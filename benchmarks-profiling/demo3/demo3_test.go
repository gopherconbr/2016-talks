package main

import (
	"io/ioutil"
	"testing"
)

func BenchmarkSineCircle(b *testing.B) {
	w := ioutil.Discard
	for i := 0; i < b.N; i++ {
		sineCircle(w, 10)
	}
}
