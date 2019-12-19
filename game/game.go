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

func (g Game) GameStart(capacity int, start chan time.Time) {
	for g.time = capacity; g.time > 0; g.time-- {
		time.Sleep(1000 * time.Millisecond)
		now := time.Now()
		start <- now
		//start <- g.time

	}
	close(start)
}
