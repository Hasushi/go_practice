package main

import (
	"io"
	"os"
)

type Rot13Reader struct {
	r io.Reader
}

func (rr *Rot13Reader) Read(p []byte) (int, error) {
	n, err := rr.r.Read(p)
	if err != nil {
		return n, err
	}
	for i := 0; i < n; i++ {
		if ('a' <= p[i] && p[i] <= 'm') || ('A' <= p[i] && p[i] <= 'M') {
			p[i] += 13
		} else {
			p[i] -= 13
		}
	}
	return n, err
}

func main() {
	rr := &Rot13Reader{os.Stdin}
	io.Copy(os.Stdout, rr)
}