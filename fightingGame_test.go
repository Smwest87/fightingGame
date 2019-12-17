package fightingGame

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	p1 = character{
		position: 1,
		speed:    1,
		hp:       1,
	}

	p2 = character{
		position: 10,
		speed:    1,
		hp:       1,
	}

	location = level{
		size: 10,
	}
)

func TestMovePlayer(t *testing.T) {
	t.Run("Success -- Move Player By + 1", func(tt *testing.T) {

		err := p1.movePlayer(&p1, &p2, &location, 1)

		assert.Nil(tt, err)
		assert.Equal(tt, 2, p1.position)
	})

	t.Run("Failure -- Outside of Level", func(tt *testing.T) {
		p1.position = 1
		err := p1.movePlayer(&p1, &p2, &location, -1)

		assert.EqualError(tt, err, "Player cannot leave the map, Player New Position : 0")
		assert.Equal(tt, 1, p1.position)
	})

	t.Run("Failure -- Players in same positin", func(tt *testing.T) {
		p2.position = 2
		err := p1.movePlayer(&p1, &p2, &location, 1)
		assert.EqualError(tt, err, "players cannot have the same position")
		assert.Equal(tt, 1, p1.position)
	})
}
