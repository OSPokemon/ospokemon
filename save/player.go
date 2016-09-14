package save

import (
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/space"
	"time"
)

const COMP_Player = "save/Player"
const EVNT_PlayerPush = "save,Player.Push"
const EVNT_PlayerPull = "save.Player.Pull"
const EVNT_PlayerDelete = "save.Player.Delete"

type Player struct {
	Username   string
	Level      uint
	Experience uint
	Money      uint
	Entity     *engine.Entity
}

func (p *Player) Id() string {
	return COMP_Player
}

func (p *Player) Update(*engine.Universe, *engine.Entity, time.Duration) {
	// TODO
}

func MakePlayer(username string) *Player {
	p := &Player{
		Username:   username,
		Level:      0,
		Experience: 0,
		Money:      0,
		Entity:     engine.MakeEntity(),
	}

	p.Entity.AddComponent(p)

	p.Entity.AddComponent(&engine.Location{
		Shape: &space.Rect{
			Anchor:    space.Point{},
			Dimension: space.Vector{},
		},
	})

	p.Entity.AddComponent(make(engine.Bindings))

	return p
}

var Players = make(map[string]*Player)
