package fighter

import (
	"time"

	"github.com/smwest87/fightingGame/actions"
)

// Position ...
type Position struct {
	X float64
	Y float64
}

// Moveset contains all possible player attacks w range, dmg info
type Moveset struct {
	Jab    map[string]float64
	Medium map[string]float64
	Heavy  map[string]float64
	Super  map[string]float64
}

// Character struct contains all player-level info
type Character struct {
	Position
	Moveset
	Defense float64
	Speed   float64
	Jump    float64
	Hp      float64
}

// Fighter interface contains all possible player actions
type Fighter interface {
	MovePlayer(playerChan PlayerChannel, xForward bool, yUp bool)

	Jab(playerChan PlayerChannel)

	Medium(playerChan PlayerChannel)

	Heavy(playerChan PlayerChannel)

	Super(playerChan PlayerChannel)

	Block(playerChan PlayerChannel)
}

// PlayerChannel stores Actions for read from game.go
type PlayerChannel chan *actions.Action

// TODO: rewrite player actions to send *actions.Action to PlayerChannel

// MovePlayer prepares an input with Movement info
func (p *Character) MovePlayer(playerChan PlayerChannel, xForward bool, yUp bool) {
	input := actions.Action{}
	input.Movement.StartTime = time.Now()

	switch xForward {
	case true:
		input.Movement.Xdist += p.Speed
	case false:
		input.Movement.Xdist -= p.Speed
	}

	switch yUp {
	case true:
		input.Movement.Ydist += p.Jump
	case false:
		break
	}

	playerChan <- &input
}

// Jab prepares an Attack input with dmg, range determined by Moveset.Jab info
func (p *Character) Jab(playerChan PlayerChannel) {
	input := actions.Action{}
	input.Attack.StartTime = time.Now()

	input.Attack.Damage = p.Moveset.Jab["damage"]
	input.Attack.Range = p.Moveset.Jab["range"]

	playerChan <- &input
}

// Medium prepares a medium attack input
func (p *Character) Medium(playerChan PlayerChannel) {
	input := actions.Action{}
	input.Attack.StartTime = time.Now()

	input.Attack.Damage = p.Moveset.Medium["damage"]
	input.Attack.Range = p.Moveset.Medium["range"]

	playerChan <- &input
}

// Heavy prepares a heavy attack input
func (p *Character) Heavy(playerChan PlayerChannel) {
	input := actions.Action{}
	input.Attack.StartTime = time.Now()

	input.Attack.Damage = p.Moveset.Heavy["damage"]
	input.Attack.Range = p.Moveset.Heavy["range"]

	playerChan <- &input
}

// Super prepares a super attack input
func (p *Character) Super(playerChan PlayerChannel) {
	input := actions.Action{}
	input.Attack.StartTime = time.Now()

	input.Attack.Damage = p.Moveset.Super["damage"]
	input.Attack.Range = p.Moveset.Super["range"]

	playerChan <- &input
}

// Block prepares a Block input
func (p *Character) Block(playerChan PlayerChannel) {
	input := actions.Action{}
	input.Block.StartTime = time.Now()

	input.Block.Defense = p.Defense

	playerChan <- &input
}

// Legacy code

// func (p *Character) movePlayer(playerOne *Character, playerTwo *Character, l *level.Level, distance int) error {

// 	if playerOne.Position.X+distance <= 0 || playerOne.Position.X+distance >= l.Size {
// 		return fmt.Errorf("Player cannot leave the map, Player New Position : %v", playerOne.Position.X+distance)
// 	}

// 	if playerOne.Position.X+distance == playerTwo.Position.X {
// 		return errors.New("players cannot have the same position")
// 	}
// 	playerOne.Position.X += distance

// 	return nil
// }

// func (p *Character) jab(opponent *Character, l *level.Level) error {
// 	if p.Position.X-opponent.Position.X == 1 || p.Position.X-opponent.Position.X == -1 {
// 		opponent.Hp -= p.Attack
// 		return nil
// 	}
// 	return errors.New("Opponent out of range")
// }

// func (p *Character) block(opponent *Character) error {
// 	if p.Position.X-opponent.Position.X == 1 || p.Position.X-opponent.Position.X == -1 {
// 		return nil
// 	}
// 	return errors.New("No effective attack to block")
// }
