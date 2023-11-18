package main

import "fmt"

func main() {
	//handle with channels
	cancelWithClose()

	fmt.Println("------------")

	//handle with context
	cancelWithContext()
}
