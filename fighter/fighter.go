package fighter

import (
	"errors"
	"fmt"

	"github.com/smwest87/fightingGame/level"
)

type position struct {
	x int
	y int
}

// Character struct contains all player-level info
type Character struct {
	position
	speed       int
	hp          int
	attack      int
	isAttacking bool
	isBlocking  bool
}

// Fighter interface contains all possible player actions
type Fighter interface {
	movePlayer(playerOne *Character, playerTwo *Character, l *level.Level, distance int) error

	jab(opponent *Character, l *level.Level) error

	block(opponent *Character) error

	endAttack()

	endBlock()
}

func (p *Character) endAttack() {
	p.isAttacking = false
}

func (p *Character) endBlock() {
	p.isBlocking = false
}

func (p *Character) movePlayer(playerOne *Character, playerTwo *Character, l *level.Level, distance int) error {

	if playerOne.position.x+distance <= 0 || playerOne.position.x+distance >= l.Size {
		return fmt.Errorf("Player cannot leave the map, Player New Position : %v", playerOne.position.x+distance)
	}

	if playerOne.position.x+distance == playerTwo.position.x {
		return errors.New("players cannot have the same position")
	}
	playerOne.position.x += distance

	return nil
}

func (p *Character) jab(opponent *Character, l *level.Level) error {
	p.isAttacking = true
	defer p.endAttack()
	if p.position.x-opponent.position.x == 1 || p.position.x-opponent.position.x == -1 {
		if !opponent.isBlocking {
			opponent.hp -= p.attack
			return nil
		}
	}
	return errors.New("Opponent out of range")
}

func (p *Character) block(opponent *Character) error {
	p.isBlocking = true
	defer p.endBlock()
	if opponent.isAttacking {
		if p.position.x-opponent.position.x == 1 || p.position.x-opponent.position.x == -1 {
			return nil
		}
	}
	return errors.New("No effective attack to block")
}
