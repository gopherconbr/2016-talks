package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//Create a context with a 1 second timeout
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	fmt.Printf("%T\n", ctx)
	fmt.Printf("%T\n", cancel)
	fmt.Printf("Err: %s\n", ctx.Err())
	//Call cancel to avoid leaking resources until timeout fires
	cancel()
	fmt.Printf("Err: %s\n", ctx.Err())
}
