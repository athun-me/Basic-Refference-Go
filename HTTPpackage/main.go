package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Make a POST request to the server
	url := "http://localhost:8080"
	payload := []byte("This is the request body for a POST request.")
	resp, err := http.Post(url, "text/plain", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Printf("Response from server: %s\n", string(body))
}
