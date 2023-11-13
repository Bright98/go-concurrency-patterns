package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go sendSignal1(ch1)
	go sendSignal2(ch2)

	result1 := <-ch1
	result2 := <-ch2
	fmt.Println("receiver got signal 1: ", result1)
	fmt.Println("receiver got signal 2: ", result2)
}
func sendSignal1(ch chan string) {
	fmt.Println("sender 1 sent signal")
	ch <- "signal 1"
}
func sendSignal2(ch chan string) {
	fmt.Println("sender 2 sent signal")
	ch <- "signal 2"
}
