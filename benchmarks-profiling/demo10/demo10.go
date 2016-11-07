package main

import (
	"io"
	"os"
	"strconv"
	"sync"
)

var hexMap [256][]byte

func init() {
	for i := 0; i <= 0xff; i++ {
		hexMap[i] = []byte(strconv.FormatInt(int64(i), 16)) // HL
	}
}

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
	buffer := pool.Get().([]byte)[:0]
	buffer = append(buffer, "{\"myformat\":"...)
	for _, b := range data {
		buffer = append(buffer, "["...)
		buffer = append(buffer, hexMap[b]...) // HL
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
			pool.Put(data)
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
