package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const URL = "http://localhost:3000"

func main() {

	fmt.Println("Server in Golang")
	PerformGetRequest(URL)
	PerformPostFormRequest(URL + "/post")

}

// ----------------- GET Request ----------------- //

func PerformGetRequest(u string) {
	fmt.Println("Performing GET request to " + u)
	fmt.Println("----------------------------------------")
	// Perform GET request here
	response, err := http.Get(u)
	checkError(err)
	defer response.Body.Close()

	fmt.Println("Response status:", response.Status)
	fmt.Println("Response size in KB:", response.ContentLength)

	fmt.Println("----------------------------------------")
	// Read the response body
	fmt.Println("Reading the response body")

	var responseString strings.Builder     // Creating a variable of type strings.Builder to store the response body
	body, err := io.ReadAll(response.Body) // Again we use io.ReadAll to read the response body
	checkError(err)                        // fmt.Println("Response body:", string(body))

	byteCount, err := responseString.Write(body)
	checkError(err)
	// Write the body to the string builder. We use Write because it returns the number of bytes written
	// and we use responseString.String() to get the string representation of the builder.
	// It is more efficient than using string concatenation or doing string(body) because
	// it doesn't create a new string every time we append to it.

	fmt.Println("Byte count: ", byteCount)
	fmt.Println("Response body:", responseString.String())
	fmt.Println("========================================")

}

// ----------------- POST Request ----------------- //
// PerformPostFormRequest performs a POST request to the specified URL with form data
func PerformPostFormRequest(u string) {
	fmt.Println("Performing POST request to " + u)
	fmt.Println("----------------------------------------")

	// Create form data using url.Values
	data := url.Values{}      // Vlaues is a type that represents form data. It is a map[string][]string type.
	data.Add("name", "Monal") // Add key-value pairs to the form data
	data.Add("age", "22")
	data.Add("job", "Web Developer")

	// Perform POST request with form data
	res, err := http.PostForm(u, data)
	checkError(err)
	defer res.Body.Close()

	fmt.Println("Response status:", res.Status)
	fmt.Println("----------------------------------------")

	// Read the response body
	fmt.Println("Reading the response body")
	content, err := io.ReadAll(res.Body) // Read the response body using io.ReadAll
	checkError(err)

	// Print the response body
	var responseString strings.Builder
	responseString.Write(content) // This will write the content to the variable responseString of type strings.Builder
	fmt.Println("Response body:", responseString.String())
	fmt.Println("========================================")
}

// checkError checks if an error occurred and panics if it did
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
