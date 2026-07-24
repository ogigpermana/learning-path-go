// File: middle/json/main.go
// Level: Middle
// Topik: JSON Encoding/Decoding
//
// Package encoding/json untuk serialisasi data Go ke JSON dan sebaliknya.
// JSON tags pada struct mengontrol nama field di JSON:
//   `json:"name"`           -> field name
//   `json:"name,omitempty"` -> hilangkan jika nilai kosong
//   `json:"-"`              -> skip field ini

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Struct Person dengan JSON tags
// Tag memberitahu encoder/decoder JSON tentang mapping
type Person struct {
	Name string `json:"name"`             // field JSON: "name"
	Age  int    `json:"age"`              // field JSON: "age"
	City string `json:"city,omitempty"`   // omitempty: hilangkan jika ""
	// Unexported field (huruf kecil) tidak akan di-serialize
	secret string
}

func main() {
	// 1. MARSHAL - Encode Go struct ke JSON string
	fmt.Println("=== Marshal (encode Go -> JSON) ===")
	p := Person{Name: "Anggi", Age: 20, City: "Jakarta"}
	data, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println("JSON:", string(data))

	// MarshalIndent untuk pretty print
	dataPretty, _ := json.MarshalIndent(p, "", "  ")
	fmt.Println("Pretty JSON:")
	fmt.Println(string(dataPretty))

	// 2. UNMARSHAL - Decode JSON ke Go struct
	fmt.Println("\n=== Unmarshal (decode JSON -> Go) ===")
	jsonStr := `{"name":"Budi","age":25}`
	var p2 Person
	err = json.Unmarshal([]byte(jsonStr), &p2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decoded struct: %+v\n", p2)

	// 3. MAP[string]interface{} - untuk JSON dinamis
	fmt.Println("\n=== Dynamic JSON with map ===")
	jsonMap := `{"name":"Citra","age":30,"hobbies":["coding","reading"]}`
	var result map[string]interface{}
	json.Unmarshal([]byte(jsonMap), &result)

	// Accessing map values need type assertion
	fmt.Println("Name:", result["name"])
	hobbies := result["hobbies"].([]interface{})
	fmt.Println("Hobbies:")
	for _, h := range hobbies {
		fmt.Println("-", h)
	}

	// 4. JSON ENCODER/DECODER - untuk file I/O
	fmt.Println("\n=== JSON Encoder/Decoder ===")
	file, err := os.Create("data.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Encoder menulis langsung ke file (lebih efisien untuk file besar)
	encoder := json.NewEncoder(file)
	encoder.Encode(p)
	fmt.Println("Data written to data.json")

	// Decoder membaca dari file
	file2, _ := os.Open("data.json")
	defer file2.Close()

	var p3 Person
	json.NewDecoder(file2).Decode(&p3)
	fmt.Printf("Read from file: %+v\n", p3)

	// Cleanup
	os.Remove("data.json")
}