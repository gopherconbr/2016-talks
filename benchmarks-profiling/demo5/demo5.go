package main

import (
	"bufio"
	"io"
	"math"
	"os"
)

func sineCircle(w io.Writer, r int) {
	buf := bufio.NewWriter(w) // HL
	defer buf.Flush()         // HL
	xScale := 2.0
	for l := 0; l <= r*2; l++ {

		alpha := math.Acos(float64(r-l) / float64(r))
		x := int(float64(r) * math.Sin(alpha) * xScale)

		padding := int(xScale*float64(r)) - x
		for j := 0; j < padding; j++ {
			buf.Write([]byte(" "))
		}

		for j := 0; j <= x*2; j++ {
			buf.Write([]byte("*"))
		}

		buf.Write([]byte("\n"))
	}
}

func main() {
	sineCircle(os.Stdout, 10)
}
