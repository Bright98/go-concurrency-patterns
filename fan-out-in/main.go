package main

import (
	"fmt"
	"sync"
)

func doTask(task int, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	result := fmt.Sprintf("task %d done", task)
	out <- result
}
func main() {
	tasks := []int{10, 11, 50, 30, 54, 60, 90, 70}
	results := make(chan string, len(tasks))
	var wg sync.WaitGroup

	//Fan Out: start multiple workers
	for _, task := range tasks {
		wg.Add(1)
		go doTask(task, results, &wg)
	}

	//Fan In: collect results and close it
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("Result:", result)
	}
}
