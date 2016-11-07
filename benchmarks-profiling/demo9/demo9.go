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

var pool = sync.Pool{ // HL
	New: func() interface{} {
		return make([]byte, 100) // HL
	},
}

func (e *encoder) encode(data []byte) {
	buffer := pool.Get().([]byte)[:0] // HL
	buffer = append(buffer, "{\"myformat\":"...)
	for _, b := range data {
		buffer = append(buffer, "["...)
		buffer = append(buffer, strconv.FormatInt(int64(b), 16)...)
		buffer = append(buffer, "]"...)
	}
	buffer = append(buffer, "}"...)
	e.ch <- buffer
}

func (e *encoder) consumer() {
	e.wg.Add(1)
	go func() {
		defer e.wg.Done()
		for data := range e.ch {
			e.w.Write(data)
			pool.Put(data) // HL
		}
	}()
}

func main() {
	ch := make(chan []byte)
	enc := encoder{ch: ch, w: os.Stdout}
	enc.consumer()
	enc.encode([]byte("is there anybody out there?"))
	close(ch)
	enc.wg.Wait()
}
