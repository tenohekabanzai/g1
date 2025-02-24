package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Post struct {
	Id      string `json:"id"`
	Content string `json:"content"`
	Creator User   `json:"creator"`
}

type User struct {
	Username string `json:"Username"`
	Email    string `json:"email"`
}

var Posts = []Post{
	{"001", "abc def ghij", User{"Abc", "abc@gmail.com"}},
	{"002", "pqr stu vwxy", User{"Def", "def@gmail.com"}},
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(Posts)
	return
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	for _, val := range Posts {
		if val.Id == id {
			json.NewEncoder(w).Encode(val)
			return
		}
	}
	return
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqbody, _ := ioutil.ReadAll(r.Body)
	var newPost Post
	_ = json.Unmarshal(reqbody, &newPost)
	Posts = append(Posts, newPost)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqbody, _ := ioutil.ReadAll(r.Body)
	var Temp []Post
	var newPost Post
	_ = json.Unmarshal(reqbody, &newPost)
	for _, val := range Posts {
		if val.Id != newPost.Id {
			Temp = append(Temp, val)
		}
	}
	Temp = append(Temp, newPost)
	Posts = Temp
	return
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var Temp []Post
	for _, val := range Posts {
		if val.Id != params["id"] {
			Temp = append(Temp, val)
		}
	}
	Posts = Temp
	return
}

func main() {
	fmt.Println("Server running on PORT 5004")
	r := mux.NewRouter()
	r.HandleFunc("/GetAll", GetAllUsers).Methods("GET")
	r.HandleFunc("/Get/{id}", GetUser).Methods("GET")
	r.HandleFunc("/add", CreatePost).Methods("POST")
	r.HandleFunc("/update", UpdatePost).Methods("PUT")
	r.HandleFunc("/delete/{id}", DeletePost).Methods("DELETE")
	http.ListenAndServe(":5004", r)
}
