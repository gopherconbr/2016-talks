package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func main() {
	myHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	}
	s := &http.Server{
		Addr:         ":8080",
		Handler:      myHandler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
