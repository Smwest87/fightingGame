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
	State     GameState
}

type GameState struct {
	timeStamp time.Time
	gameClock int
}

func (g Game) GameStart(capacity int, start chan GameState) {
	for g.State.gameClock = capacity; g.State.gameClock >= 0; g.State.gameClock-- {
		g.State.timeStamp = time.Now()
		start <- g.State
		time.Sleep(1000 * time.Millisecond)

	}
	close(start)
}
