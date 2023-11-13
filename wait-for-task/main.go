package main

import (
	"fmt"
	"sync"
	"time"
)

//wait for task done with wait group from sync package

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go task(&wg)

	wg.Wait()
	fmt.Println("task done")
}
func task(wg *sync.WaitGroup) {
	fmt.Println("task is doing")
	time.Sleep(time.Second)
	wg.Done()
}
