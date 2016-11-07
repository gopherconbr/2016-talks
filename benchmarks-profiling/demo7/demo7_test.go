package main

import (
	"bufio"
	"os"
	"testing"
)

func BenchmarkSineCircle(b *testing.B) {
	out, _ := os.Open("/dev/null")
	bufOut := bufio.NewWriter(out) // HL
	for i := 0; i < b.N; i++ {
		sineCircle(bufOut, 10)
	}
	bufOut.Flush() // HL
}
