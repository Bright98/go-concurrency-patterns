package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	result := make(chan any, 2)
	var wg sync.WaitGroup

	username := getUsername()

	wg.Add(2)
	go getLikes(username, result, &wg)
	go getBio(username, result, &wg)

	wg.Wait()
	close(result)

	for res := range result {
		fmt.Println(res)
	}

	fmt.Println("It takes: ", time.Since(start))
}

func getUsername() string {
	time.Sleep(time.Millisecond * 150)
	return "My username"
}
func getLikes(username string, res chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 50)
	res <- 20
	wg.Done()
}
func getBio(username string, res chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	res <- "My name is test"
	wg.Done()
}
