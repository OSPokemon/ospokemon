package engine

import (
	"github.com/ospokemon/ospokemon/util"
	"math"
	"time"
)

const EVNT_StatUpdate = "engine.Stat.Update"

type Stat struct {
	Value     float64
	Max       float64
	TempMax   float64
	Regen     float64
	TempRegen float64
}

func (s *Stat) Update(u *Universe, e *Entity, c *Chart, n string, d time.Duration) {
	if s.Value < 1 {
		return
	}

	s.Value += math.Min(s.TempMax-s.Value, s.TempRegen)
	util.Event.Fire(EVNT_StatUpdate, u, e, c, n, s)
	u.Fire(EVNT_StatUpdate, u, e, c, n, s)
	e.Fire(EVNT_StatUpdate, u, e, c, n, s)
}
