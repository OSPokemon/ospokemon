package comp

import (
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/util"
	"time"
)

const PLAYER = "Player"

type Player save.Player

func init() {
	util.Event.On(save.EVNT_PlayersNew, func(args ...interface{}) {
		playersplayernew(args[0].(*save.Player))
	})
	util.Event.On(save.EVNT_PlayersGet, func(args ...interface{}) {
		playersplayerget(args[0].(*save.Player))
	})
}

func (p *Player) Id() string {
	return PLAYER
}

func (p *Player) Snapshot() map[string]interface{} {
	return map[string]interface{}{
		"username":   p.Username,
		"level":      p.Level,
		"experience": p.Experience,
		"money":      p.Money,
	}
}

func (p *Player) Update(*engine.Universe, *engine.Entity, time.Duration) {
	// TODO
}

func playersplayernew(p *save.Player) {
	comp := Player(*p)
	p.Entity.AddComponent(&comp)
}

func playersplayerget(p *save.Player) {
	comp := Player(*p)
	p.Entity.AddComponent(&comp)
}
