package game

import (
	"time"

	"github.com/smwest87/fightingGame/fighter"
	"github.com/smwest87/fightingGame/level"
)

type Game struct {
	time      int
	location  level.Level
	playerOne fighter.Character
	playerTwo fighter.Character
}

func (g Game) GameStart(start chan int) {
	for g.time = 60; g.time > 0; g.time-- {
		time.Sleep(1000 * time.Millisecond)
		//fmt.Println(time.Now())
		start <- g.time

	}
}
