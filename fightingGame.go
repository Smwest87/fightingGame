package fightingGame

import (
	"errors"
	"fmt"
)

type position struct {
	x int
	y int
}

type character struct {
	position
	speed  int
	hp     int
	attack int
}

type level struct {
	size int
}

func (p *character) movePlayer(playerOne *character, playerTwo *character, l *level, distance int) error {

	if playerOne.position.x+distance <= 0 || playerOne.position.x+distance >= l.size {
		return fmt.Errorf("Player cannot leave the map, Player New Position : %v", playerOne.position.x+distance)
	}

	if playerOne.position.x+distance == playerTwo.position.x {
		return errors.New("players cannot have the same position")
	}
	playerOne.position.x += distance

	return nil
}

func (p *character) jab(opponent *character, l *level) error {
	if p.position.x-opponent.position.x == 1 || p.position.x-opponent.position.x == -1 {
		opponent.hp -= p.attack
		return nil
	}
	return errors.New("Opponent out of range")
}
