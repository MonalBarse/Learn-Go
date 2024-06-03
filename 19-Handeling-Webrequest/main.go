package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Web Requests in Golang")
	response, err := http.Get("https://the-internet.herokuapp.com/abtest")
	checkNilError(err)

	defer response.Body.Close() // Ensure the body is closed after we are done with it
	// fmt.Println("Response: ", response)
	fmt.Println("Status: ", response.Status)
	fmt.Println("StatusCode: ", response.StatusCode)
	// fmt.Println("Header: ", response.Header)
	fmt.Println("ContentLength: ", response.ContentLength/(1024*1024), "MB")
	fmt.Println("Uncompressed: ", response.Uncompressed, "Description: This field is set to true if the response was sent compressed but was decompressed by the http package automatically. This feature helps in handling gzip or deflate compression transparently.")
	fmt.Println("Proto: ", response.Proto)           // Protp is HTTP protocol version eg. HTTP/1.1 or HTTP/2.0
	fmt.Println("ProtoMajor: ", response.ProtoMajor) //The major version number of the HTTP protocol used by the server, for example, 1 for HTTP/1.1 or 2 for HTTP/2.
	fmt.Println("ProtoMinor: ", response.ProtoMinor) // The minor version number of the HTTP protocol used by the server, for example, 1 for HTTP/1.1.
	fmt.Println("Close: ", response.Close)
	// Indicates whether the server asked to close the connection after the response is sent.
	fmt.Println("Uncompressed: ", response.Uncompressed)

	// Read and print the body of the response
	body, err := io.ReadAll(response.Body)
	checkNilError(err)
	fmt.Println("Body: ", string(body))

}

func checkNilError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
