package main

import "fmt"

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, _ := http.NewRequest("GET", "http://example.com", nil)
	req.WithContext(ctx)

	client := &http.Client{}
	client.Do(req)
}
