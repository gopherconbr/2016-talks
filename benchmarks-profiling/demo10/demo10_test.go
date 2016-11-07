package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func BenchmarkEncodePrint(b *testing.B) {
	ch := make(chan []byte)                   // OMIT
	enc := encoder{ch: ch, w: ioutil.Discard} // OMIT
	// ...
	enc.consumer()
	data := bytes.Repeat([]byte("*"), 1024)
	for i := 0; i < b.N; i++ {
		enc.encode(data)
	}
	// ...
	close(ch)     // OMIT
	enc.wg.Wait() // OMIT
}
