package main

import (
	"context"
	"fmt"
	"time"
)

func expensiveOp(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("Still calculating nothing")
	}
}

func main() {
	calcDone := make(chan struct{})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		expensiveOp(ctx)
		calcDone <- struct{}{}
	}()

	select {
	case <-calcDone:
		fmt.Println("Calculation done")
	case <-ctx.Done():
		fmt.Println("Calculation Timeout")
	}
}
