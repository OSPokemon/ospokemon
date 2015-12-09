package effectscripts

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/world"
	"time"
)

type stateeffect byte

var State stateeffect

func (s stateeffect) New(name string, state uint8, duration time.Duration) *world.Effect {
	return &world.Effect{
		Name:     name,
		Priority: world.PRIOstate,
		Data: map[string]interface{}{
			"state": state,
		},
		Script:   State.Script,
		Duration: duration,
	}
}

func (e stateeffect) Script(effect *world.Effect, entity world.Entity, now time.Time) {

	data, ok := effect.Data["state"].(uint8)
	if !ok {
		log.WithFields(log.Fields{
			"Entity": entity,
		}).Error("effectscripts.State invalid data supplied")
		return
	}

	mortal, ok := entity.(world.Mortality)
	if !ok {
		log.WithFields(log.Fields{
			"Entity": entity,
		}).Error("effectscripts.State invalid target")
		return
	}

	switch data {
	case world.CTRLimmune:
	case world.CTRLstasis:
	case world.CTRLcloak:
		mortal.SetControl(mortal.Control() | data)
		return
	case world.CTRLstun:
	case world.CTRLroot:
		if world.IsProtected(mortal) {
			effect.Duration = 0
			return
		}

		mortal.SetControl(mortal.Control() | data)
		return
	}
}
