package main

import (
	"bytes"
	"io"
	"math"
	"net/http"
	_ "net/http/pprof" // HL
)

func sineCircle(w io.Writer, r int) {
	xScale := 2.0
	for l := 0; l <= r*2; l++ {

		alpha := math.Acos(float64(r-l) / float64(r))
		x := int(float64(r) * math.Sin(alpha) * xScale)

		padding := int(xScale*float64(r)) - x
		w.Write(bytes.Repeat([]byte(" "), padding)) // HL
		w.Write(bytes.Repeat([]byte("*"), x*2))     // HL
		w.Write([]byte("\n"))
	}
}

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sineCircle(w, 10)
	}))
	http.ListenAndServe("0.0.0.0:10101", nil) // HL
}
