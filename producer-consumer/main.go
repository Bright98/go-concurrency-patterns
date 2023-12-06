package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(ch chan<- int) {
	for {
		data := rand.Intn(100)
		ch <- data
		fmt.Println("produced: ", data)
		time.Sleep(time.Millisecond * 100)
	}
}
func consumer1(ch <-chan int) {
	for data := range ch {
		fmt.Println("consume 1: ", data)
	}
	time.Sleep(time.Millisecond * 300)
}
func consumer2(ch <-chan int) {
	for data := range ch {
		fmt.Println("consume 2: ", data)
	}
	time.Sleep(time.Millisecond * 150)
}
func main() {
	ch := make(chan int)
	go producer(ch)
	go consumer1(ch)
	go consumer2(ch)

	time.Sleep(time.Second * 1)
}
