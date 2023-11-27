package main

import (
	"fmt"
	"time"
)

type Semaphore interface {
	Acquire()
	Release()
}
type semaphore struct {
	ch chan struct{}
}

func New(capacity int) Semaphore {
	return &semaphore{ch: make(chan struct{}, capacity)}
}
func (s *semaphore) Acquire() {
	s.ch <- struct{}{}
}
func (s *semaphore) Release() {
	<-s.ch
}
func longProcess(index int) {
	fmt.Println(time.Now().Format("15:04:05"), "Running task: ", index)
	time.Sleep(time.Second * 2)
}
func main() {
	sem := New(3)
	done := make(chan bool, 1)
	total := 10

	for i := 1; i <= total; i++ {
		sem.Acquire()
		go func(index int) {
			defer sem.Release()
			longProcess(index)

			if index == total {
				done <- true
			}
		}(i)
	}

	<-done
}
