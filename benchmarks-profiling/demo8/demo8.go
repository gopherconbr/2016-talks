package main

import (
	"io"
	"os"
	"strconv"
	"sync"
)

type encoder struct {
	w  io.Writer
	ch chan []byte
	wg sync.WaitGroup
}

func (e *encoder) encode(data []byte) {
	var buffer []byte
	buffer = append(buffer, "{\"myformat\":"...)
	for _, b := range data {
		buffer = append(buffer, "["...)
		buffer = append(buffer, strconv.FormatInt(int64(b), 16)...)
		buffer = append(buffer, "]"...)
	}
	buffer = append(buffer, "}"...)
	e.ch <- buffer // HL
}

func (e *encoder) consumer() {
	e.wg.Add(1)
	go func() {
		defer e.wg.Done()
		for data := range e.ch { // HL
			e.w.Write(data)
		}
	}()
}

func main() {
	ch := make(chan []byte)
	enc := encoder{ch: ch, w: os.Stdout}              // HL
	enc.consumer()                                    // HL
	enc.encode([]byte("is there anybody out there?")) // HL
	close(ch)
	enc.wg.Wait()
}
