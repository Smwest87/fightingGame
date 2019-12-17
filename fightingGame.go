package fightingGame

import (
	"errors"
	"fmt"
)

type character struct {
	position int
	speed    int
	hp       int
}
type level struct {
	size int
}

func (p *character) movePlayer(playerOne *character, playerTwo *character, l *level, distance int) error {

	if playerOne.position+distance <= 0 || playerOne.position+distance >= l.size {
		return fmt.Errorf("Player cannot leave the map, Player New Position : %v", playerOne.position+distance)
	}

	if playerOne.position+distance == playerTwo.position {
		return errors.New("players cannot have the same position")
	}
	playerOne.position += distance

	return nil
}
