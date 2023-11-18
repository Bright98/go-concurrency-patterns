package main

import (
	"fmt"
	"sync"
)

func doTask(id int, tasks []int, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		result := fmt.Sprintf("task %d is doing by worker %d", task, id)
		out <- result
	}
}
func main() {
	tasks := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	workers := 3
	results := make(chan string, len(tasks))
	var wg sync.WaitGroup

	//create worker pool
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go doTask(i, tasks, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	//get results
	for result := range results {
		fmt.Println("Result:", result)
	}
}
