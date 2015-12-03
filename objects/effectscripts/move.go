package effectscripts

import (
	"github.com/ospokemon/ospokemon/world"
	// "github.com/ospokemon/ospokemon/world/update"
	"log"
	"time"
)

type moveeffect byte

var Move moveeffect

func (e moveeffect) New(vector *world.Vector, duration time.Duration) *world.Effect {
	return &world.Effect{
		Name:     "Move",
		Priority: PRIOstandard,
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
