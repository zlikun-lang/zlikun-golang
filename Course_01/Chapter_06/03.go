package main

import (
	"fmt"
	"io"
)

type ReadWriter interface {
	io.Reader
	io.Writer
}

type Example struct {
	name string
}

func (e Example) Read(p []byte) (n int, err error) {
	// Mock
	fmt.Printf("read: %s\n", string(p))
	return len(p), nil
}

func (e Example) Write(p []byte) (n int, err error) {
	// Mock
	fmt.Printf("write: %s\n", string(p))
	return len(p), nil
}

func x(rw ReadWriter) {
	_, err1 := rw.Read([]byte("10001"))
	if err1 != nil {

	}

	_, err2 := rw.Write([]byte("10002"))
	if err2 != nil {

	}
}

func main() {
	var example Example
	x(example)
}
