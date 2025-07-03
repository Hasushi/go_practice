package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

type LogWriter struct {
	bw *bufio.Writer
}

func (lw *LogWriter) Write(p []byte) (int, error) {

	ts := time.Now().Format(time.RFC3339) + " "
	if _, err := lw.bw.WriteString(ts); err != nil {
		fmt.Print("failed to write to file")
		return 0, err
	}

	if _, err := lw.bw.Write(p); err != nil {
		fmt.Print("failed to write to file")
		return 0, err
	}
	if err := lw.bw.Flush(); err != nil {
		return 0, err
	}

	return len(p), nil

}

func main() {
	wf, err := os.Create("test.txt")
	if err != nil {
		panic("failed to open the file")
	}
	defer wf.Close()

	lw := &LogWriter{bufio.NewWriter(wf)}

	tee := io.TeeReader(os.Stdin, lw)
	
	if _, err = io.Copy(os.Stdout, tee); err != nil {
		panic("panic")
	}
	
}