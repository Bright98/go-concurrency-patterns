package main

import "fmt"

func work(ch chan struct{}) {
	fmt.Println("running...")
	for {
		select {
		case <-ch:
			fmt.Println("canceled !")
			return
		}
	}
}
func main() {
	ch := make(chan struct{})
	go work(ch)

	//cancel channel
	fmt.Println("cancelling...")
	ch <- struct{}{}
}
