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
	jsonData, err := json.Marshal(payload)
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
