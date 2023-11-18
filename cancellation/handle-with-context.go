package main

import (
	"context"
	"fmt"
	"time"
)

func cancelWithContext() {
	duration := time.Second
	root := context.Background()
	ctx, cancel := context.WithTimeout(root, duration)
	defer cancel()

	ch := make(chan struct{}, 1)
	go task(ch, duration)

	select {
	case <-ch:
		fmt.Println("work complete")
	case <-ctx.Done():
		fmt.Println("work cancelled")
	}
}
func task(ch chan struct{}, duration time.Duration) {
	//work cancel
	//sleep := duration * 2
	//work complete
	sleep := duration - time.Millisecond*500
	time.Sleep(sleep)
	ch <- struct{}{}
}
