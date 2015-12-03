package effectscripts

import (
	"github.com/ospokemon/ospokemon/world"
	"github.com/ospokemon/ospokemon/world/update"
	"time"
)

type moveeffect byte

var Move moveeffect

func (e moveeffect) New(vector *world.Vector, now time.Time, duration time.Duration) *world.Effect {
	return &world.Effect{
		Name:     "Move",
		Priority: world.PRIOstandard,
		Data:     vector,
		Script:   Move.Script,
		Start:    now,
		Duration: duration,
	}
}

func (h *moveeffect) Script(effect *world.Effect, entity world.Entity, now time.Time) {

	if entity.Controls().State&world.CTRLPstuck > 0 {
		return
	}

	vector := effect.Data.(*world.Vector)
	update.MoveEntity(entity, vector)
	return
}
