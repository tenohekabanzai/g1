package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://chaiwithcode.com/"

func main() {
	fmt.Println("Web Requests")

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	//fmt.Println(resp)

	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))

}
