package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type Player struct {
	health int32
}

func NewPlayer() *Player {
	return &Player{
		health: 100,
	}
}
func (p *Player) getHealth() int {
	return int(atomic.LoadInt32(&p.health))
}
func (p *Player) takeDamage(value int) {
	health := p.getHealth()
	atomic.StoreInt32(&p.health, int32(health-value))
}
func startUILoop(p *Player) {
	//reading data
	ticker := time.NewTicker(time.Second)
	for {
		fmt.Println("player health is: ", p.getHealth())
		<-ticker.C
	}
}
func startGameLoop(p *Player) {
	//writing data
	ticker := time.NewTicker(time.Millisecond * 300)
	for {
		p.takeDamage(rand.Intn(30))
		if p.health <= 0 {
			fmt.Println("GAME OVER")
			break
		}
		<-ticker.C
	}
}

func main() {
	player := NewPlayer()
	go startUILoop(player)
	startGameLoop(player)
}
