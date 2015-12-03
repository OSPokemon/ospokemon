package effectscripts

import (
	"github.com/ospokemon/ospokemon/world"
	"log"
	"time"
)

type stateeffect byte

var State stateeffect

func (s stateeffect) New(state uint8, now time.Time, duration time.Duration) *world.Effect {
	return &world.Effect{
		Name:     "StateMod",
		Priority: world.PRIOstate,
		Data:     state,
		Script:   State.Script,
		Start:    now,
		Duration: duration,
	}
}

func (e stateeffect) Script(effect *world.Effect, entity world.Entity, now time.Time) {

	data, ok := effect.Data.(uint8)
	if !ok {
		log.Printf("effectscripts.State invalid data supplied: %v\n", effect.Data)
		return
	}

	isprotected := entity.Controls().State&world.CTRLPprotected < 1

	switch data {
	case world.CTRLimmune:
	case world.CTRLstasis:
	case world.CTRLcloak:
		entity.Controls().State |= data
		return
	case world.CTRLstun:
	case world.CTRLroot:
		if isprotected {
			effect.Duration = 0
			return
		}

		entity.Controls().State |= data
		return
	}
}
