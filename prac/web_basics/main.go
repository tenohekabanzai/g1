package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web Requests")
	performGetReq()
	fmt.Println("---------------------------------------")
	performPostReq()
	fmt.Println("---------------------------------------")
	performPostFormReq()
}

func performGetReq() {
	_url := "http://localhost:5004/get"
	resp, _ := http.Get(_url)
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

func performPostReq() {
	_url := "http://localhost:5004/post"
	// mimics json
	x := strings.NewReader(`[{"username":"Abc","age":21,"email":"xyz@example.com"},{"username":"Abc","age":21,"email":"xyz@example.com"}]`)

	resp, _ := http.Post(_url, "application/json", x)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

func performPostFormReq() {

	_url := "http://localhost:5004/postform"
	// x := strings.NewReader(`[{"username":"Abc","age":21,"email":"xyz@example.com"},{"username":"Abc","age":21,"email":"xyz@example.com"}]`)

	// mimics form
	x := url.Values{}
	x.Add("name", "Abc")
	x.Add("age", "20")
	x.Add("email", "Abc@gmail.com")

	resp, _ := http.PostForm(_url, x)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}
