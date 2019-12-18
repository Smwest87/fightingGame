package fighter

import (
	"testing"

	"github.com/smwest87/fightingGame/level"
	"github.com/stretchr/testify/assert"
)

var (
	p1 = Character{
		position: position{
			x: 1,
			y: 0,
		},
		speed:  1,
		hp:     1,
		attack: 1,
	}

	p2 = Character{
		position: position{
			x: 10,
			y: 0,
		},
		speed:  1,
		hp:     1,
		attack: 1,
	}

	location = level.Level{
		Size: 10,
	}
)

func TestMovePlayer(t *testing.T) {
	t.Run("Success -- Move Player By + 1", func(tt *testing.T) {

		err := p1.movePlayer(&p1, &p2, &location, 1)

		assert.Nil(tt, err)
		assert.Equal(tt, 2, p1.position.x)
	})

	t.Run("Failure -- Outside of Level - 0", func(tt *testing.T) {
		p1.position.x = 1
		err := p1.movePlayer(&p1, &p2, &location, -1)

		assert.EqualError(tt, err, "Player cannot leave the map, Player New Position : 0")
		assert.Equal(tt, 1, p1.position.x)
	})

	t.Run("Failure -- Outside of Level - 11", func(tt *testing.T) {
		err := p1.movePlayer(&p2, &p1, &location, 1)

		assert.EqualError(tt, err, "Player cannot leave the map, Player New Position : 11")
		assert.Equal(tt, 10, p2.position.x)
	})

	t.Run("Failure -- Players in same position", func(tt *testing.T) {
		p2.position.x = 2
		err := p1.movePlayer(&p1, &p2, &location, 1)
		assert.EqualError(tt, err, "players cannot have the same position")
		assert.Equal(tt, 1, p1.position.x)
	})
}

func TestJab(t *testing.T) {
	t.Run("Jab in Range", func(tt *testing.T) {
		err := p1.jab(&p2, &location)
		assert.Nil(tt, err)
		assert.Equal(tt, 0, p2.hp)
	})

	t.Run("Jab out of Range", func(tt *testing.T) {
		p2.hp = 1
		p2.position.x = 10
		err := p1.jab(&p2, &location)
		assert.EqualError(tt, err, "Opponent out of range")
		assert.Equal(tt, 1, p2.hp)
	})
}

func TestBlock(t *testing.T) {
	t.Run("Successfully block incoming attack", func(tt *testing.T) {
		// move player 2 adjacent to player 1 and execute jab
		p2.position.x = 2
		p1StartHp := p1.hp
		_ = p2.jab(&p1, &location)

		// player 1 blocks incoming attack
		blockErr := p1.block(&p2)
		p1EndHp := p1.hp
		assert.Nil(tt, blockErr)
		assert.Equal(tt, p1StartHp, p1EndHp)
	})
}
