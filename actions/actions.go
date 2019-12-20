package action

import (
	"time"
)

// Action struct contains all info for a player input
type Action struct {
	Movement *Movement
	Attack   *Attack
	Block    *Block
}

// Attack struct represents data for attack action
type Attack struct {
	Damage   float64
	Range    float64
	Startup  time.Duration
	Active   time.Duration
	Recovery time.Duration
}

// Block contains all block related info
type Block struct {
	Defense  float64
	Startup  time.Duration
	Active   time.Duration
	Recovery time.Duration
}

// Movement contains all movement input info
type Movement struct {
	Xdist float64
	Ydist float64
	Speed float64
}
