package main

import (
	"fmt"
	"time"
)

func announce(message string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Println(message)
}

func main() {
	go announce("Hello Gophers", 2*time.Second)

	fmt.Println("Gophercon Brazil 2016")

	time.Sleep(3 * time.Second)
}
