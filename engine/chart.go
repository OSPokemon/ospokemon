package engine

import (
	"github.com/ospokemon/ospokemon/util"
	"time"
)

const COMP_Chart = "engine/Chart"
const EVNT_ChartUpdate = "engine.Chart.Update"

type Chart struct {
	Stats map[string]*Stat
	Buffs []*Buff
}

func (c *Chart) Id() string {
	return COMP_Chart
}

func (c *Chart) Update(u *Universe, e *Entity, d time.Duration) {
	for _, s := range c.Stats {
		s.TempMax = s.Max
		s.TempRegen = s.Regen
	}

	for i := 0; i < len(c.Buffs); i++ {
		if !c.Buffs[i].Update(u, e, c, d) {
			copy(c.Buffs[i:], c.Buffs[i+1:])
			c.Buffs[len(c.Buffs)-1] = nil
			c.Buffs = c.Buffs[:len(c.Buffs)-1]
			i--
		}
	}

	for n, s := range c.Stats {
		s.Update(u, e, c, n, d)
	}

	util.Event.Fire(EVNT_ChartUpdate, u, e, c)
	u.Fire(EVNT_ChartUpdate, u, e, c)
	e.Fire(EVNT_ChartUpdate, u, e, c)
}
