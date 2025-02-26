package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Location struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Country   string  `json:"country"`
	State     string  `json:"state"`
}

var apikey string

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error in parsing .env file")
		return
	}
	apikey = os.Getenv("api_key")
	// fmt.Println(apikey)
	r := mux.NewRouter()
	r.HandleFunc("/weather", callApi).Methods("POST")
	fmt.Println("server running at port 5004")
	http.ListenAndServe(":5004", r)

}

func callApi(w http.ResponseWriter, r *http.Request) {

	var city map[string]string
	json.NewDecoder(r.Body).Decode(&city)
	c := city["city"]
	// fmt.Println(c)
	geo_url := fmt.Sprintf("https://geocoding-by-api-ninjas.p.rapidapi.com/v1/geocoding?city=%v", c)
	req2, _ := http.NewRequest("GET", geo_url, nil)
	req2.Header.Add("x-rapidapi-key", apikey)
	req2.Header.Add("x-rapidapi-host", "geocoding-by-api-ninjas.p.rapidapi.com")

	res2, _ := http.DefaultClient.Do(req2)
	defer res2.Body.Close()
	data, _ := ioutil.ReadAll(res2.Body)
	// str_data := string(data)

	var locations []Location
	err := json.Unmarshal(data, &locations)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	Lat := fmt.Sprintf("%f", locations[0].Latitude)
	Lon := fmt.Sprintf("%f", locations[0].Longitude)

	url := fmt.Sprintf("https://weather-by-api-ninjas.p.rapidapi.com/v1/weather?lat=%s&lon=%s", Lat, Lon)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("x-rapidapi-key", apikey)
	req.Header.Add("x-rapidapi-host", "weather-by-api-ninjas.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
