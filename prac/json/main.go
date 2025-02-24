package main

import (
	"encoding/json"
	"fmt"
)

type Emp struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func encodeJson() {
	list := []Emp{
		{"Abc", 20, "abc@gmail.com"},
		{"Def", 20, "def@gmail.com"},
		{"Ghi", 20, "ghi@gmail.com"},
	}

	finalJson, _ := json.Marshal(list)
	//fmt.Printf("%T\n", finalJson)
	fmt.Println(string(finalJson))
}

func decodeJson() {

	json1 := []byte(`[
		{"name":"Abc","age":20,"email":"abc@gmail.com"},
		{"name":"Def","age":20,"email":"def@gmail.com"},
		{"name":"Ghi","age":20,"email":"ghi@gmail.com"}
	]`)
	var emp []Emp
	_ = json.Unmarshal(json1, &emp)
	fmt.Println(emp)

	var emp2 []map[string]interface{}
	_ = json.Unmarshal(json1, &emp2)

	for _, val := range emp2 {
		// fmt.Println(val)
		for key, v := range val {
			fmt.Println(key, v)
		}
	}
}

func main() {
	fmt.Println("JSON")
	// encodeJson()
	decodeJson()
}
