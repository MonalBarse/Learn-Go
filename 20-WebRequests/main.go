package main

import (
	"fmt"
	"net/url"
)

const URL = "https://the-internet.herokuapp.com:3040/abtest?name=Monal&language=go&hello=world#test"

func main() {

	fmt.Println("Handeling URL in Golang")
	fmt.Println("URL: ", URL)

	// URL Parsing
	// The URL package provides a URL type that represents a URL. The URL type has methods to parse a URL string, construct a URL string, and query the URL. The URL type is used to represent a URL and provides methods to parse a URL string, construct a URL string, and query the URL.

	// The URL type has the following fields:
	// Scheme: The URL scheme, such as http, https, ftp, etc.
	// Opaque: The encoded opaque data.
	// User: The username.
	// Password: The password.
	// Host: The host, including the port number.
	// Path: The path.
	// RawPath: The encoded path hint (Go 1.5 and later).
	// RawQuery: The encoded query values, without '?'. For example, "k=v&k=v".
	// Fragment: The fragment for references, without '#'.
	// The URL type has the following methods:
	// Parse: Parse parses a URL string and returns a URL type.
	// String: String returns the URL string.
	// Query: Query returns the query values.

	fmt.Println("URL Parsing")

	// Parse a URL string
	result, err := url.Parse(URL)
	checkNilError(err)

	fmt.Println("Scheme: ", result.Scheme)
	fmt.Println("Opaque: ", result.Opaque)
	fmt.Println("User: ", result.User)
	fmt.Println("Host: ", result.Host)
	fmt.Println("Port: ", result.Port())
	fmt.Println("Path: ", result.Path)
	fmt.Println("RawQuery: ", result.RawQuery)
	fmt.Println("Fragment: ", result.Fragment)
	fmt.Println("String: ", result.String())
	fmt.Println("Query: ", result.Query())
	fmt.Println("We can loop through the query values")
	// Or we can loop throguh the query values
	for key, value := range result.Query() {
		fmt.Println(key, ":", value)
	}

	partsOfURL := &url.URL{ // The partsOfURL is a pointer to the URL type. If we had parts of the URL and wanted to construct a URL string, we could use the URL type to construct the URL string.
		Scheme:   "https",
		Host:     "the-internet.herokuapp.com",
		Path:     "/bcTest",
		RawQuery: "name=Goku&language=js",
	}
	anotherUrl := partsOfURL.String()
	fmt.Println("Another URL: ", anotherUrl)

}

func checkNilError(err error) {

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
