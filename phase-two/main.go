package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func main(){
	// create a new http client
	client := &http.Client{
		Timeout: time.Second * 10, // timeout for every request
	}

	// make a POST request
	url := "https://example.com/api"
	method := "POST"
	payload := strings.NewReader(`{"key1":"val1", "key2":"val2"}`)

	req, err := http.NewRequest(method,url,payload)
	if err != nil {
		fmt.Println("Error creating request", err)
		return
	}

	// Set headers
	req.Header.Add("Content-Type","application/json")
	req.Header.Add("Authorization", "Bearer TOKEN")

	// exec request
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer response.Body.Close()

	// read and print response
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println(string(body))
	// https://medium.com/@hakimnaufal/simple-go-http-requests-get-and-response-handling-b6fd83d62a5e
}