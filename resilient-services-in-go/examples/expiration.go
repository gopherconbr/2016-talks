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
	defer cancel()

	fmt.Println("Waiting for timeout")
	<-ctx.Done()
	fmt.Println("Done, Reason: ", ctx.Err())
}
