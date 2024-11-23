package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Define the URL to send the request to
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// Send a GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the response
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))
}
