package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
)

const api = "https://localhost:8080"
const tryCount = 3

var loading = []rune{
	'⠏',
	'⠛',
	'⠹',
	'⠼',
	'⠶',
	'⠧',
}

var (
	maxConnections = runtime.NumCPU()
	payloadSize = []int{
		1562500,  // 1.5625MB
		6250000,  // 6.25MB
		12500000, // 12.5MB
		26214400, // 25MB
	}
)

func main() {
	ctx := context.Background()

	var (
		lastDown string
		downBytes int64
		lastUp string
		upBytes int64
	)
	fmt.Println("Starting download test...")

	var i int
	err := DownloadTest(ctx, func(result *Lap) error {
		lastDown = result.String()
		downBytes = result.Bytes
		fmt.Printf("%c%s, size: %d ↓ - %s bps, size: %d ↑\r", loading[i%len(loading)], lastDown, result.Bytes, "", 0)
		i++
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	err = UploadTest(ctx, func(result *Lap) error {
		lastUp = result.String()
		upBytes = result.Bytes
		fmt.Printf("%c%s, size: %d ↓ - %s bps, size: %d ↑\r", loading[i%len(loading)], lastUp, result.Bytes, "", 0)
		i++
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%c%s, size: %d ↓ - %s bps, size: %d ↑\r", loading[i%len(loading)], lastDown, downBytes, lastUp, upBytes)
}