/* package main

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
} */

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Response struct {
	IsProfane  bool    `json:"isProfane"`
	Score      float64 `json:"score"`
	FlaggedFor string  `json:"flaggedFor,omitempty"`
}

func main() {
	fmt.Println("Sending POST request to Slift API")

	// Define the API endpoint
	apiURL := "https://slift-api.monalbarse.workers.dev/"

	// Prepare the request payload as a map because to convert it to JSON we need to marshal it (which is the process of converting a Go object to JSON)
	payload := map[string]string{
		"message": "What da fuck",
	}

	// Convert payload to JSON
	jsonData, err := json.Marshal(payload) // we do json.Marshal becuase
	checkError(err)

	// Create a new POST request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	checkError(err)

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	checkError(err)
	defer resp.Body.Close()

	// Read and parse the response body
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, resp.Body); err != nil {
		checkError(err)
	}

	// Print the raw response body for debugging
	fmt.Println("Raw Response Body:", buf.String())

	// Decode JSON response
	var response Response
	err = json.Unmarshal(buf.Bytes(), &response)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		os.Exit(1)
	}

	// Print decoded response
	fmt.Println("Is Profane:", response.IsProfane)
	fmt.Println("Score:", response.Score)
	if response.IsProfane {
		fmt.Println("Flagged For:", response.FlaggedFor)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
