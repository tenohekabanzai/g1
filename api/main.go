package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// slice mimics a DB
var courses = []Course{
	{CourseId: "101", CourseName: "Go Basics", CoursePrice: 499, Author: &Author{Fullname: "John Doe", Website: "johndoe.com"}},
	{CourseId: "102", CourseName: "Advanced Go", CoursePrice: 799, Author: &Author{Fullname: "Jane Smith", Website: "janesmith.com"}},
}

func (c *Course) isEmpty() bool { return c.CourseId == "" || c.CourseName == "" }

func main() {

	fmt.Println("REST API")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome)
	r.HandleFunc("/course", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{Id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{Id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{Id}", deleteCourse).Methods("DELETE")
	http.ListenAndServe(":5003", r)

}

// controllers
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello from REST API</h1>"))
	return
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
	return
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	//get id from params
	params := mux.Vars(r)

	for _, val := range courses {
		if val.CourseId == params["Id"] {
			json.NewEncoder(w).Encode(val)
			return
		}
	}

	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		fmt.Println("Hello1")
		json.NewEncoder(w).Encode("Incomplete details provided")
	}
	var c Course
	_ = json.NewDecoder(r.Body).Decode(&c)

	c.CourseId = strconv.Itoa(rand.Intn(1000))

	courses = append(courses, c)
	json.NewEncoder(w).Encode(c)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, val := range courses {
		if val.CourseId == params["Id"] {
			courses = append(courses[:idx], courses[idx+1:]...)
			var c Course
			_ = json.NewDecoder(r.Body).Decode(&c)
			// fmt.Println(c)
			c.CourseId = params["Id"]
			courses = append(courses, c)
			json.NewEncoder(w).Encode(c)
			return
		}
	}
	return
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, val := range courses {
		if val.CourseId == params["Id"] {
			courses = append(courses[:idx], courses[idx+1:]...)
			json.NewEncoder(w).Encode("Deleted course with provided id")
			return
		}
	}
	return
}
