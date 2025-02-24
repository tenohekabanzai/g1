package main

import (
	"fmt"
	"net/url"
)

const _url = "https://lco.dev:3000/test?user=abc&type=paid"

func main() {
	fmt.Println("Hello")

	resp, _ := url.Parse(_url)

	fmt.Println(resp)
	fmt.Println(resp.Scheme)
	fmt.Println(resp.Host)
	fmt.Println(resp.Path)
	q := resp.Query()
	for key, val := range q {
		fmt.Println(key, val)
	}

	new_url := url.URL{
		Scheme:   "https",
		Host:     "lco.dev:4001",
		Path:     "/learn",
		RawQuery: "course=NextJS&level=advanced",
	}
	_url2 := new_url.String()
	fmt.Println(_url2)
	_url2_parsed, _ := url.Parse(_url2)
	fmt.Println(_url2_parsed.Query())
}
