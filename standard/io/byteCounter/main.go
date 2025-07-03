package main

import (
	"fmt"
	"io"
	"os"
)

type ByteCounter struct {
	N int64
}

func (bc *ByteCounter) Write(p []byte) (int, error) {
	n := len(p)
	bc.N += int64(n)
	return n, nil
}

func main() {
	bc := &ByteCounter{}
	io.Copy(bc, os.Stdin)
	fmt.Printf("Total bytes is %d", bc.N)
}