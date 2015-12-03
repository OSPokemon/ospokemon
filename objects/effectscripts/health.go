package effectscripts

import (
	"github.com/ospokemon/ospokemon/world"
	"log"
	"time"
)

type healtheffect byte

var Health stateeffect

func (e healtheffect) New(power int, now time.Time, duration time.Duration) *world.Effect {
	return &world.Effect{
		Name:     "HealthMod",
		Priority: world.PRIOstandard,
		Data:     power,
		Script:   Health.Script,
		Start:    now,
		Duration: duration,
	}
}

func (h *healtheffect) Script(effect *world.Effect, entity world.Entity, now time.Time) {

	power, ok := effect.Data.(int)
	if !ok {
		log.Printf("effectscripts.State invalid data supplied: %v\n", effect.Data)
		return
	}
	healthy, ok := entity.(world.Healthy)
	if !ok {
		log.Printf("effectscripts.State invalid target supplied: %v\n", entity)
		return
	}

	isprotected := entity.Controls().State&world.CTRLPprotected < 1

	if power < 0 && isprotected {
		return
	}

	health := healthy.Health() + power

	if health > healthy.MaxHealth() {
		power = healthy.MaxHealth() - healthy.Health()
	} else if health < 0 {
		power = -healthy.Health()
	}

	healthy.SetHealth(healthy.Health() + power)
}
