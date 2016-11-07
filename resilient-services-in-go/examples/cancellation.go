package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	go func() {
		cancel()
	}()

	fmt.Println("Not waiting for timeout")
	<-ctx.Done()
	fmt.Println("Done, Reason: ", ctx.Err())
}
