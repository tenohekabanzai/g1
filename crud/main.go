package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string   `json:"id"`
	Isbn     string   `json:"isbn"`
	Title    string   `json:"title"`
	Director Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {

	movies = append(movies, Movie{Id: "1", Isbn: "894734", Title: "Movie One", Director: Director{Firstname: "Abc", Lastname: "Def"}})
	movies = append(movies, Movie{Id: "2", Isbn: "222234", Title: "Movie Two", Director: Director{Firstname: "Pqr", Lastname: "Def"}})
	movies = append(movies, Movie{Id: "3", Isbn: "891134", Title: "Movie Three", Director: Director{Firstname: "Ghi", Lastname: "Jkl"}})
	movies = append(movies, Movie{Id: "4", Isbn: "798734", Title: "Movie Four", Director: Director{Firstname: "Abc", Lastname: "Def"}})

	fmt.Println(movies)

	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")
	fmt.Println("Server Running at PORT 5004")
	http.ListenAndServe(":5004", r)
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
	return
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	for _, val := range movies {
		if val.Id == id {
			json.NewEncoder(w).Encode(val)
			return
		}
	}
	return
}

func createMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m Movie
	_ = json.NewDecoder(r.Body).Decode(&m)

	movies = append(movies, m)
	json.NewEncoder(w).Encode(movies)
	return
}
func updateMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	for i, val := range movies {
		if val.Id == id {
			movies = append(movies[:i], movies[i+1:]...)
			var m Movie
			_ = json.NewDecoder(r.Body).Decode(&m)
			movies = append(movies, m)
		}
	}

	json.NewEncoder(w).Encode(movies)
	return
}
func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	for i, val := range movies {
		if val.Id == id {
			movies = append(movies[:i], movies[i+1:]...)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
	json.NewEncoder(w).Encode(movies)
	return
}
