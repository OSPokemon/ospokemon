package ospokemon

import (
	"ospokemon.com/log"
	"time"
)

const PARTspawner = "spawner"

type Spawner struct {
	Child *Entity
	Speed time.Duration
	Timer *time.Duration
}

func MakeSpawner() *Spawner {
	return &Spawner{}
}

func (s *Spawner) Part() string {
	return PARTspawner
}

func (spawner *Spawner) Update(u *Universe, d time.Duration) {
	if spawner.Child.Id > 0 {
		return
	}

	if spawner.Timer == nil {
		timer := spawner.Speed
		spawner.Timer = &timer
		log.Add("Universe", spawner.Child.UniverseId).Add("Timer", timer).Info("spawner: reset")
	}

	if *spawner.Timer < d {
		spawner.Timer = nil
		u.Add(spawner.Child)
		log.Add("Universe", spawner.Child.UniverseId).Add("Entity", spawner.Child.Id).Info("spawner: activate")
	} else {
		*spawner.Timer = *spawner.Timer - d
	}
}
