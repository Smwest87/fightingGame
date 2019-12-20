package game

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/smwest87/fightingGame/fighter"
	"github.com/smwest87/fightingGame/level"
)

var headLine string = "              0               0      "
var bodyLine string = "             /|\\             /|\\    "
var weapon string = "             (o-\"=         =\"-o)    "
var legs string = "-------------/ \\-------------/ \\--------"

type Game struct {
	time      int
	location  level.Level
	playerOne fighter.Character
	playerTwo fighter.Character
	State     GameState
	Frame     string
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

	g.playerOne.Hp = 1
	g.playerOne.Position.X = 1

	g.playerTwo.Hp = 1
	g.playerTwo.Position.X = 10
	fmt.Println("---------------------Ready?---------------------")
	fmt.Println("---------------------Fight!---------------------")

	clock := make(chan GameState, 5)

	// uncomment when these are ready for use
	// p1Chan := make(chan fighter.PlayerChannel)
	// p2Chan := make(chan fighter.PlayerChannel)

	go g.GameStart(cap(clock), clock)

	for stamp := range clock {
		if g.playerOne.Hp > 0 || g.playerTwo.Hp > 0 || stamp.GameClock != 0 {
			print("\033[H\033[2J")
			g.State.GameClock = stamp.GameClock
			println(g.StartFrame())
		}
	}

	fmt.Println("\r\n\r\n---------------------Game Over!---------------------")

}

func (g Game) StartFrame() string {

	renderP1 := "|"
	p1Position := g.playerOne.Position.X
	p1Index := int(math.Floor(p1Position))

	renderP2 := "|"
	p2Position := g.playerTwo.Position.X
	p2Index := int(math.Floor(p2Position))

	ground := []string{"-", "-", "-", "-", "-", "-", "-", "-", "-", "-"}

	ground[p1Index-1] = renderP1
	ground[p2Index-1] = renderP2

	return fmt.Sprintf("     %v     \r\n%v", g.State.GameClock, strings.Join(ground, ""))

	//return fmt.Sprintf("                     %v       \r\n\r\n\r\n\r\n\r\n%v\r\n%v\r\n%v\r\n%v", g.State.GameClock, headLine, bodyLine, weapon, legs)
}
