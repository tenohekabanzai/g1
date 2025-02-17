package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Firstname string
	Lastname  string
	Age       int
	Email     string
}

func encodeJSON() {

	pslice := []Person{
		{"Abc", "Def", 11, "abc@gmail.com"},
		{"Ghi", "Jkl", 13, "ghi@gmail.com"},
		{"Mno", "pqr", 11, "mno@gmail.com"},
	}

	jsonEnc, _ := json.Marshal(pslice)
	jsonEncIndent, _ := json.MarshalIndent(pslice, "", "\t")
	fmt.Println(string(jsonEnc))
	fmt.Println(string(jsonEncIndent))
}

func decodeJSON() {

	// decoding into a struct
	jsonfromWeb := []byte(`[{"firstname":"Abc","lastname":"def","age":32,"email":"abc@gmail.com"},{"firstname":"Ghi","lastname":"Jkl","age":23,"email":"ghi@gmail.com"}]`)

	var p1 []Person
	json.Unmarshal(jsonfromWeb, &p1)
	fmt.Println(p1)

	//decoding into a map
	jsonfromWeb2 := []byte(`{"firstname":"Abc","lastname":"def","age":32,"email":"abc@gmail.com"}`)
	var mp map[string]interface{}
	json.Unmarshal(jsonfromWeb2, &mp)
	fmt.Println(mp)

}

func main() {
	fmt.Println("JSON")
	encodeJSON()

	decodeJSON()
}
