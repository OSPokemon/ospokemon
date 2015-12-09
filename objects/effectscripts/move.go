package effectscripts

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/update"
	"github.com/ospokemon/ospokemon/world"
	"time"
)

type moveeffect byte

var Move moveeffect

func (e moveeffect) New(vector *world.Vector, duration time.Duration) *world.Effect {
	return &world.Effect{
		Name:     "Move",
		Priority: world.PRIOstandard,
		Data: map[string]interface{}{
			"Vector": vector,
		},
		Script:   Move.Script,
		Duration: duration,
	}
}

func (h *moveeffect) Script(effect *world.Effect, entity world.Entity, now time.Time) {
	mortal, ok := entity.(world.Mortality)
	if !ok {
		log.WithFields(log.Fields{
			"target": entity,
		}).Error("effectscripts.Move invalid target supplied")
	}

	if world.IsStuck(mortal) {
		return
	}

	vector, ok := effect.Data["Vector"].(*world.Vector)
	if !ok {
		log.WithFields(log.Fields{
			"data": effect.Data,
		}).Error("effectscripts.Move invalid data supplied")
	}

	update.MoveEntity(entity, vector)
	return
}
