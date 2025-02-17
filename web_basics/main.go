package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func performGetReq() {
	const url = "http://localhost:5004/get"
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	r, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(r))
}

func performPostReq() {
	const url = "http://localhost:5004/post"
	reqBody := strings.NewReader(`{ "name":"Abc","email":"abc@gmail.com"}`)
	resp, _ := http.Post(url, "application/json", reqBody)
	defer resp.Body.Close()
	r, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(r))
}

func performPostForm() {
	const _url = "http://localhost:5004/postform"
	data := url.Values{}
	data.Add("firstname", "abc")
	data.Add("laststname", "def")
	data.Add("email", "abc@gmail.com")

	resp, _ := http.PostForm(_url, data)
	defer resp.Body.Close()
	r, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(r))
}

func main() {
	performGetReq()
	performPostReq()
	performPostForm()
}
