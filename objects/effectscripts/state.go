package effectscripts

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/world"
	"time"
)

type stateeffect byte

var State stateeffect

func (s stateeffect) New(state uint8, duration time.Duration) *world.Effect {
	return &world.Effect{
		Name:     "StateMod",
		Priority: world.PRIOstate,
		Data: map[string]string{
			"state": string(state),
		},
		Script:   State.Script,
		Duration: duration,
	}
}

func (e stateeffect) Script(effect *world.Effect, entity world.Entity, now time.Time) {

	data := uint8([]byte(effect.Data["state"])[0])

	mortal, ok := entity.(world.Mortality)
	if !ok {
		log.WithFields(log.Fields{
			"entity": entity,
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
