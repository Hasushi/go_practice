package main

import (
	// "fmt"
	"compress/gzip"
	"encoding/base64"
	"io"
	"os"
)

type CountWriter struct {
	Count int64
}

func (cw *CountWriter) Write(p []byte) (int, error) {
	n := len(p)
	cw.Count += int64(n)
	return n, nil
}

type UpperReader struct {
	r io.Reader
}

func (ur *UpperReader) Read(p []byte) (int, error) {
	n, err := ur.r.Read(p)
	if err != nil {
		return n, err
	}
	for i := 0; i < n; i++ {
		if 'a' <= p[i] && p[i] <= 'z' {
			p[i] = p[i] - 'a' + 'A'
		}
	}
	return n, nil
}

func gzip_base64_decode() {

	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b64 := base64.NewEncoder(base64.StdEncoding, f)
	defer b64.Close()

	gz := gzip.NewWriter(b64)
	defer gz.Close()

	io.Copy(gz, os.Stdin)
}

func main() {
	gzip_base64_decode()
}