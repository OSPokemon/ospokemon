package engine

import (
	"github.com/ospokemon/ospokemon/util"
	"time"
)

const COMP_Chart = "ospokemon/engine/Chart"
const EVNT_ChartUpdate = "ospokemon/engine/Chart.Update"

type Chart struct {
	Stats map[string]*Stat
	Buffs []Buff
}

func (c *Chart) Id() string {
	return COMP_Chart
}

func (c *Chart) Update(u *Universe, e *Entity, d time.Duration) {
	for _, s := range c.Stats {
		s.TempMax = s.Max
		s.TempRegen = s.Regen
	}

	for _, b := range c.Buffs {
		b.Update(u, e, c, d)
	}

	for n, s := range c.Stats {
		s.Update(u, e, c, n, d)
	}

	util.Event.Fire(EVNT_ChartUpdate, u, e, c)
	u.Fire(EVNT_ChartUpdate, u, e, c)
	e.Fire(EVNT_ChartUpdate, u, e, c)
}
