package main

import (
	"fmt"
	"time"
)

func getResult(ch chan int) {
	for p := range ch {
		fmt.Println("received signal: ", p)
	}
}
func main() {
	const capacity = 10
	ch := make(chan int, capacity)

	go getResult(ch)

	const work = 20
	for w := 0; w < work; w++ {
		select {
		case ch <- w:
			fmt.Println("sent signal: ", w)
		default:
			fmt.Println("dropper data: ", w)
		}
	}

	close(ch)
	fmt.Println("shutdown signal")

	time.Sleep(time.Second)
}
