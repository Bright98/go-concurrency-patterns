package main

import "testing"

// go test . -> pass
// but
// go test . --race failed because of data racing
// answer: using Mutex

func TestGame(t *testing.T) {
	player := NewPlayer()
	go startUILoop(player)
	startGameLoop(player)
}
