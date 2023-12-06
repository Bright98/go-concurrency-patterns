package main

import (
	"fmt"
)

var ch = make(chan int)
var done = make(chan struct{})

func producer(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	done <- struct{}{}
}
func consumer(ch <-chan int) {
	for data := range ch {
		fmt.Println("consume: ", data)
	}
}
func main() {
	go producer(ch)
	go consumer(ch)
	<-done
}
