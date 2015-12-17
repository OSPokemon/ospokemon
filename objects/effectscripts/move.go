package effectscripts

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/physics"
	"github.com/ospokemon/ospokemon/update"
	"github.com/ospokemon/ospokemon/world"
	"time"
)

func NewMoveEffect(name string, vector physics.Vector) *world.Effect {
	return &world.Effect{
		Name:     name,
		Priority: world.PRIOstandard,
		Data: map[string]interface{}{
			"Vector": vector,
		},
		Script:   MoveScript,
		Duration: 0,
	}
}

func MoveScript(effect *world.Effect, entity world.Entity, now time.Time) {
	mortal, ok := entity.(world.Mortality)
	if !ok {
		log.WithFields(log.Fields{
			"target": entity,
		}).Error("effectscripts.Move invalid target supplied")
	}

	if world.IsStuck(mortal) {
		return
	}

	vector, ok := effect.Data["Vector"].(physics.Vector)
	if !ok {
		log.WithFields(log.Fields{
			"data": effect.Data,
		}).Error("effectscripts.Move invalid data supplied")
	}

	update.MoveEntity(entity, vector)
	return
}
