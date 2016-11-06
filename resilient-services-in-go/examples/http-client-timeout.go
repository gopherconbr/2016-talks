package main

import "fmt"

func main() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	client.Get("http://example.com")
}
