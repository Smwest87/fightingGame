package main

import (
	"fmt"
	"time"

	"github.com/smwest87/fightingGame/game"
)

func main() {
	var g game.Game

	start := make(chan time.Time, 10)

	g.GameStart(cap(start), start)

	for stamp := range start {
		fmt.Println(stamp)
	}

}
