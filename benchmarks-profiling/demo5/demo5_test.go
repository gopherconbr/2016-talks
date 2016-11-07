package main

import (
	"os"
	"testing"
)

func BenchmarkSineCircle(b *testing.B) {
	out, _ := os.Open("/dev/null")
	for i := 0; i < b.N; i++ {
		sineCircle(out, 10)
	}
}
