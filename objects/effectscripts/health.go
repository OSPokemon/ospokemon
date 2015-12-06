package effectscripts

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/world"
	"strconv"
	"time"
)

type healtheffect byte

var Health healtheffect

func (e healtheffect) New(power int, duration time.Duration) *world.Effect {
	return &world.Effect{
		Name:     "HealthMod",
		Priority: world.PRIOstandard,
		Data: map[string]string{
			"power": strconv.Itoa(power),
		},
		Script:   Health.Script,
		Duration: duration,
	}
}

func (h *healtheffect) Script(effect *world.Effect, entity world.Entity, now time.Time) {

	power, err := strconv.Atoi(effect.Data["power"])
	if err != nil {
		log.WithFields(log.Fields{
			"data": effect.Data,
		}).Error("effectscripts.Health invalid data supplied")
		return
	}
	mortal, ok := entity.(world.Mortality)
	if !ok {
		log.WithFields(log.Fields{
			"target": entity,
		}).Error("effectscripts.Health invalid target supplied")
	}

	if power < 0 && world.IsProtected(mortal) {
		return
	}

	health := mortal.Stats()["health"].Value() + power

	if health > mortal.Stats()["health"].MaxValue() {
		power = mortal.Stats()["health"].MaxValue() - mortal.Stats()["health"].Value()
	} else if health < 0 {
		power = -mortal.Stats()["health"].Value()
	}

	mortal.Stats()["health"].SetValue(mortal.Stats()["health"].Value() + power)
}
