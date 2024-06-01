package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")
	// in JS we have objects, in Python we have dictionaries, in Golang we have maps
	// maps are key value pairs which are unordered and unindexed in nature

	// syntax of map : var map_name map[key_type]value_type
	// var languages map[string]string // If we initilaize like this, it will be nil map the print will give us : map[]

	var spokenLanguages map[string]string
	fmt.Println("If we initilaize an empty map we get: ", spokenLanguages)
	fmt.Println("----------------------------")

	// We can also initialize a map using make function similar to slices

	languages := make(map[string]string)
	fmt.Println("The value of languages map before putting any value inside: ", languages) // map[]

	// Adding values to the map
	languages["en"] = "English"
	languages["es"] = "Spanish"
	languages["fr"] = "French"
	languages["hi"] = "Hindi"
	languages["ja"] = "Japanese"
	languages["ko"] = "Korean"

	fmt.Println("The value of languages map after putting values inside: ", languages) // map[en:English es:Spanish fr:French hi:Hindi ja:Japanese ko:Korean]

	// Accessing the values from the map
	fmt.Println("The value of key 'en' in the languages map is: ", languages["en"]) // English

	// Deleting the values from the map
	delete(languages, "ko")
	fmt.Println("The value of languages map after deleting the value of key 'ko': ", languages) // map[en:English es:Spanish fr:French hi:Hindi ja:Japanese]
	fmt.Println("----------------------------")

	// Loops in maps
	fmt.Println("Looping through the map")

	for key, value := range languages {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}
	fmt.Println("----------------------------")

	// Checking if a key exists in the map
	_, exists := languages["ko"]
	fmt.Println("Does the key 'ko' exists in the languages map: ", exists) // false

	value, exists := languages["en"]
	fmt.Printf("Key: %v, Value: %v, Exists: %v\n", "en", value, exists) // Key: en, Value: English, Exists: true
	fmt.Println("----------------------------")

}
