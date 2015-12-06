package effectscripts

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/world"
	"github.com/ospokemon/ospokemon/world/update"
	"strconv"
	"time"
)

type moveeffect byte

var Move moveeffect

func (e moveeffect) New(vector *world.Vector, duration time.Duration) *world.Effect {
	return &world.Effect{
		Name:     "Move",
		Priority: world.PRIOstandard,
		Data:     map[string]string{
		// "DX" : strconv.FormatFloat(vector.DX, fmt, prec, 64)
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

	DX, err := strconv.ParseFloat(effect.Data["DX"], 64)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("effectscripts.Move invalid data supplied")
	}

	DY, err := strconv.ParseFloat(effect.Data["DY"], 64)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("effectscripts.Move invalid data supplied")
	}

	vector := &world.Vector{DX: DX, DY: DY}
	update.MoveEntity(entity, vector)
	return
}
