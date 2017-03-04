package game

import (
	"github.com/ospokemon/ospokemon/part"
	"time"
)

const DEFAULT_BAG_SIZE = 10

type Player struct {
	Username   string
	Level      uint
	Experience uint
	Money      uint
	Class      uint
	BagSize    uint
	part.Parts
}

var Players = make(map[string]*Player)

func MakePlayer(username string) *Player {
	p := &Player{
		Username: username,
		BagSize:  DEFAULT_BAG_SIZE,
		Parts:    make(part.Parts),
	}

	return p
}

func (p *Player) Part() string {
	return part.Player
}

func (p *Player) Update(u *Universe, e *Entity, d time.Duration) {
}
