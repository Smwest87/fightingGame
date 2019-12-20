package main

import (
	"fmt"

	"github.com/smwest87/fightingGame/game"
)

func main() {
	var g game.Game

	start := make(chan game.GameState, 60)

	go g.GameStart(cap(start), start)

	for stamp := range start {
		fmt.Println(stamp)
	}

}
