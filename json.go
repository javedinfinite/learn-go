package main

import (
	"encoding/json"
	"fmt"
)

// JSON uses lowercase keyname in general while struct uses pascal case. 
// So, if json has name and in struct it is Name  then the tag with name will match with json data and mapping will happen. Key will be Name only
// Only exported fields will be encoded/decoded in JSON. Fields must start with capital letters to be exported.

type Person1 struct {
	Name string `json:"full_name"` // you can control key name i.e. you can provide the key name starting with capital letter inorder to make it exported member while the tag can start with small letter as in json
	Age int `json:"age,omitempty"` // if age is zero it won't be included in the json
	Height float64
	Weight float64 `json:"-"` // this field will be ignored as json is -
	origin string // since not exported, so it will be ignored
}

func PrintJson() {

	// JSON (JavaScript Object Notation) is a lightweight format for data exchange
	// Go offers built-in support for JSON encoding and decoding, including to and from built-in and custom data types.
	fmt.Println("***************** Printing JSON*****************")
	
	jsonData := []byte(`{"full_name":"Zenia", "Age":24, "Height": 5.4, "Weight": 61.5, "origin":"India"}`) // converts a string into a slice of bytes, most Go APIs that work with JSON expect bytes, not plain strings. the standard library (like encoding/json) is designed around byte slices.
	
	marshal1, _ := json.Marshal(3.4) // converting a Go value to JSON format, here numer stays number in JSON
	marshal2, _ := json.Marshal([]string{"apple", "banana"}) // Go slice will get converted to JSON array
	marshal3, _ := json.Marshal(map[string]int{"apple":3, "banana":4}) // Go map becomes JSON object
	fmt.Println("marshal1, marshal2, marshal3 : ", string(marshal1), string(marshal2),string(marshal3)) // Converts the byte slice → string and then print
	
	var p Person1
	err := json.Unmarshal(jsonData, &p) // parse the slice of byte of json data to struct type and store it in p
	if err!=nil {
			fmt.Println("Failed to pasrse", err)
			return
		}
	fmt.Println("parsed data : ", p,p.Name,  p.Age)

	// What if we don't know the structure of the data to be parsed?
	// you can parse it into a generic type instead of a struct.

	var result map[string]interface{} // here string is the key type and  interface means the value type (means “any type”)
	err2 := json.Unmarshal(jsonData, &result)
	if err2!=nil {
		fmt.Println("Failed to parse unknown", err2)
	}
	fmt.Println("pasrsed unknown data : ", result, result["full_name"], result["Height"].(float64)) // In order to use the values in the decoded map, we’ll need to convert them to their appropriate type. For example here we convert the value in Height to the expected float64 type.
	
	
}