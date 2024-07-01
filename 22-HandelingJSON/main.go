package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string   `json:"username"`
	Age      int      `json:"age"`
	Email    string   `json:"contact"`
	Password string   `json:"-"`
	Adresses []string `json:"addresses"`
}

func main() {
	fmt.Println("Let's study about JSON encoding!")

	fmt.Println("WE will study how to create a JSON data from a Go object")
	// The Go Object we are using is a slice of User struct we could have used any other object too like a map etc
	decodeJSON()

}

func encodeJSON() {

	// JSON encoding is the process of converting a Go object to JSON format
	// The encoding/json package provides functions to encode Go objects to JSON and decode JSON to Go objects

	appUsers := []User{ // appUsers are of type []User (slice)
		{"Goku", 30, "goku@js.dev", "password456", []string{"789 Pine St", "101 Oak St"}},
		{"Vegeta", 40, "vegeta@py.dev", "password789", []string{"121 Maple St", "141 Birch St"}},
		{"Bulma", 35, "bulma@rb.dev", "password123", []string{"161 Elm St", "112 Cedar St"}},
		{"Piccolo", 150, "piccolo@db.dev", "password000", []string{"201 Walnut St", "221 Planet Namek"}},
	}

	// Let's package this data as JSON
	// The Marshal function in the encoding/json package is used to encode Go objects to JSON
	jsonData, err := json.MarshalIndent(appUsers, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	fmt.Println(string(jsonData))

}

func decodeJSON() {
	// JSON decoding is the process of converting JSON data to Go objects
	dummyJSON := []byte(`
  [
  	{"name": "Goku", "age": 30, "email": "goku@js.dev", "password": "password456", "addresses": ["789 Pine St", "101 Oak St"]},
		{"name": "Vegeta", "age": 40, "email": "vegeta@py.dev", "password": "password789", "addresses": ["121 Maple St", "141 Birch St"]},
		{"name": "Bulma", "age": 35, "email": "bulma@rb.dev", "password": "password123", "addresses": ["161 Elm St", "112 Cedar St"]},
		{"name": "Piccolo", "age": 150, "email": "piccolo@db.dev", "password": "password000", "addresses": ["201 Walnut St", "221 Planet Namek"]}
  ]
  `)
	isValid := json.Valid(dummyJSON)
	if !isValid {
		fmt.Println("The JSON is not valid")
		return
	}
	var appUsers []User
	err := json.Unmarshal([]byte(dummyJSON), &appUsers) // The Unmarshal function gives us the decoded JSON data in the appUsers slice
	// it gives out error if the JSON is not valid
	if err != nil {
		fmt.Println("Error decoding JSON")
		return
	}
	fmt.Println("----------------------------")
	fmt.Printf("%#v\n", appUsers)

	// ----------------For online data where you don't know the structure of the JSON data---------------- //
	// var result map[string]interface{}
	/* We can use the map[string]interface{} we can do this as we know the key is always a string and the value can be anything a string, int, float, bool, array, object etc */
	// err = json.Unmarshal([]byte(onlineData), &result)
	// handleErr(err)

}
