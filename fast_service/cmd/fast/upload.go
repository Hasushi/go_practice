package main

import (
	"context"
	"net/url"
	"time"

	"golang.org/x/sync/errgroup"
)

var UploadTimeout = 15 * time.Second

const uploadURL = api + "/upload"

type IntervalCallback func(current *Lap) error

func UploadTest(ctx context.Context, cb IntervalCallback) error {
	ctx, cancel := context.WithTimeout(ctx, UploadTimeout)
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
					time.Sleep(250 * time.Millisecond)
			}
			eg.Go(func() error {
				defer func() { <-semaphore }()
				if err := r.upload(ctx, uploadURL, size); err != nil {
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



func errorCheck(err error) error {
	if err == context.Canceled || err == context.DeadlineExceeded {
		return nil
	}
	if v, ok := err.(*url.Error); ok {
		err = v.Err
		return errorCheck(err)
	}
	return err
}