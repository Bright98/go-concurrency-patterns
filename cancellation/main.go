package main

import (
	"fmt"
	"time"
)

func main() {
	cancel := make(chan struct{})

	go handleRunning(cancel)

	time.Sleep(time.Second * 3)
	close(cancel)

	time.Sleep(time.Second)
}
func handleRunning(cancel chan struct{}) {
	for {
		select {
		case <-cancel:
			fmt.Println("cancelled")
			return
		default:
			fmt.Println("running...")
			time.Sleep(time.Second)
		}
	}
}
