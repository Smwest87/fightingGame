package fighter

import (
	"errors"
	"fmt"

	"github.com/smwest87/fightingGame/level"
)

type Position struct {
	X int
	Y int
}

// Character struct contains all player-level info
type Character struct {
	Position
	Speed  int
	Hp     int
	Attack int
}

// Fighter interface contains all possible player actions
type Fighter interface {
	movePlayer(playerOne *Character, playerTwo *Character, l *level.Level, distance int) error

	jab(opponent *Character, l *level.Level) error

	block(opponent *Character) error

	endAttack()

	endBlock()
}

func (p *Character) movePlayer(playerOne *Character, playerTwo *Character, l *level.Level, distance int) error {

	if playerOne.Position.X+distance <= 0 || playerOne.Position.X+distance >= l.Size {
		return fmt.Errorf("Player cannot leave the map, Player New Position : %v", playerOne.Position.X+distance)
	}

	if playerOne.Position.X+distance == playerTwo.Position.X {
		return errors.New("players cannot have the same position")
	}
	playerOne.Position.X += distance

	return nil
}

func (p *Character) jab(opponent *Character, l *level.Level) error {
	if p.Position.X-opponent.Position.X == 1 || p.Position.X-opponent.Position.X == -1 {
		opponent.Hp -= p.Attack
		return nil
	}
	return errors.New("Opponent out of range")
}

func (p *Character) block(opponent *Character) error {
	if p.Position.X-opponent.Position.X == 1 || p.Position.X-opponent.Position.X == -1 {
		return nil
	}
	return errors.New("No effective attack to block")
}
