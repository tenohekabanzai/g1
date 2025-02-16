package main

import (
	"fmt"
	"net/url"
)

const _url = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=kjsnkv"

func main() {
	fmt.Println("Handling URLS")
	fmt.Println(_url)

	res, _ := url.Parse(_url)

	fmt.Println(res.Scheme)
	fmt.Println(res.Host)
	fmt.Println(res.Path)
	fmt.Println(res.Port)
	// fmt.Println(res.Query())

	qparams := res.Query()
	fmt.Println(qparams)
	fmt.Println(qparams["coursename"])
	for _, val := range qparams {
		fmt.Println(val)
	}

	xyz := url.URL{
		Scheme:   "https",
		Host:     "lco.dev:3001",
		Path:     "/practice",
		RawQuery: "user=abc&type=Paid",
	}
	_url2 := xyz.String()

	fmt.Println(_url2)

}
