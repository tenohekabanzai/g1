package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.chaicode.com/"

func main() {
	fmt.Println("Web Reqs")

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}
