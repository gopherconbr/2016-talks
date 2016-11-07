package main

import (
	"io"
	"math"
	"os"
	"runtime/pprof"
)

func sineCircle(w io.Writer, r int) {
	xScale := 2.0
	for l := 0; l <= r*2; l++ {

		alpha := math.Acos(float64(r-l) / float64(r))
		x := int(float64(r) * math.Sin(alpha) * xScale)

		padding := int(xScale*float64(r)) - x
		for j := 0; j < padding; j++ {
			w.Write([]byte(" ")) // HL
		}

		for j := 0; j <= x*2; j++ {
			w.Write([]byte("*")) // HL
		}

		w.Write([]byte("\n")) // HL
	}
}

func main() {
	out, _ := os.Create("cpu.pprof")
	pprof.StartCPUProfile(out) // HL
	defer out.Close()
	defer pprof.StopCPUProfile() // HL

	for i := 0; i < 5000; i++ { // Arbitrary large number of iterations
		sineCircle(os.Stdout, 10)
	}
}
