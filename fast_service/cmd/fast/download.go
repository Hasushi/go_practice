package main

import (
	"context"
	"time"

	"golang.org/x/sync/errgroup"
)

var DownloadTimeout = 15 * time.Second

const downloadURL = api + "/download"

func DownloadTest(ctx context.Context, cb IntervalCallback) error {
	ctx, cancel := context.WithTimeout(ctx, DownloadTimeout)
	defer cancel()
	eg, ctx := errgroup.WithContext(ctx)

	r := newRecorder(time.Now(), maxConnections)

	go func(){
		for {
			select {
			case lap := <-r.Lap():
				cb(&lap)
			case <-ctx.Done():
				return
			}
		}
	}()

	semaphore := make(chan struct{}, maxConnections)
LOOP:
	for i := 0; i < tryCount; i++ {
		for _, size := range payloadSize {
			select {
				case <-ctx.Done():
					break LOOP
				case semaphore <- struct{}{}: 
					// 枠が空いている場合、1つ確保する
					// 250ms待機する
					time.Sleep(250 * time.Millisecond)
			}
			eg.Go(func() error {
				defer func() { <-semaphore }()
				if err := r.download(ctx, downloadURL, size); err != nil {
					return err
				}
				return nil
			})
		}
	}

	select {
	case <-ctx.Done():
	case semaphore <- struct{}{}:
		cancel()
	}
	return errorCheck(eg.Wait())
}