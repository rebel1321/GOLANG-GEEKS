package main

import (
    "fmt"
    "net/http"
)

// Function to make an HTTP GET request
func getRequest(url string) (*http.Response, error) {
    response, err := http.Get(url)
    return response, err
}

func main() {
    response, err := getRequest("http://example.com")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Response Status:", response.Status)
    }
}