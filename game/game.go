package game

import (
	"fmt"
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
	TimeStamp time.Time
	GameClock int
}

func (g Game) GameStart(capacity int, start chan GameState) {
	for g.State.GameClock = capacity; g.State.GameClock >= 0; g.State.GameClock-- {
		g.State.TimeStamp = time.Now()
		start <- g.State
		time.Sleep(1000 * time.Millisecond)

	}
	close(start)
}

func (g Game) RunGame() {

	start := make(chan GameState, 60)

	fmt.Println("---------------------Ready?---------------------")
	fmt.Println("---------------------Fight!---------------------")

	g.playerOne.Hp = 1
	g.playerTwo.Hp = 1
	go g.GameStart(cap(start), start)

	for stamp := range start {
		if g.playerOne.Hp > 0 || g.playerTwo.Hp > 0 || stamp.GameClock != 0 {
			fmt.Println(stamp.GameClock)
		}
	}

	fmt.Println("---------------------Game Over!---------------------")

}
