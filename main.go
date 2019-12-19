package main

import (
	"fmt"

	"github.com/smwest87/fightingGame/game"
)

func main() {
	var g game.Game

	start := make(chan int)

	go g.GameStart(start)

	timer := <-start
	fmt.Println(timer)
}
