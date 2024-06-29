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
	// PerformGetRequest(URL)
	PerformPostFormRequest(URL + "/post")

}

func PerformPostFormRequest(u string) {
	fmt.Println("Performing POST request to " + u)
	// Form data

	data := url.Values{}
	data.Add("name", "Monal")
	data.Add("age", "22")
	data.Add("job", "Web Developer")

	res, err := http.PostForm(u, data)
	checkError(err)

	content, err := io.ReadAll(res.Body)
	checkError(err)
	fmt.Println("Response body:", string(content))

}

func PerformGetRequest(u string) {
	fmt.Println("Performing GET request to " + u)

	// Perform GET request here

	response, err := http.Get(u)
	checkError(err)
	defer response.Body.Close()

	fmt.Println("Response status:", response.Status)
	// Get the body of the response
	fmt.Println("Response size in KB:", response.ContentLength)

	var responseString strings.Builder
	body, err := io.ReadAll(response.Body)
	checkError(err)
	// fmt.Println("Response body:", string(body))
	byteCount, err := responseString.Write(body) // Write the body to the string builder. We use Write because it returns the number of bytes written and we use responseString.String() to get the string representation of the builder. It is more efficient than using string concatenation or doing string(body) because it doesn't create a new string every time we append to it.
	checkError(err)
	fmt.Println("Byte count: ", byteCount)
	fmt.Println("Response body:", responseString.String())

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
