package engine

import (
	"github.com/ospokemon/ospokemon/util"
	"time"
)

const EVNT_BuffUpdate = "engine.Buff.Update"

type Buff struct {
	Name       string
	BonusMax   float64
	BonusRegen float64
	time.Duration
}

func (b *Buff) Update(u *Universe, e *Entity, c *Chart, d time.Duration) bool {
	stat := c.Stats[b.Name]
	stat.TempMax += b.BonusMax
	stat.TempRegen += b.BonusRegen
	b.Duration -= d

	util.Event.Fire(EVNT_BuffUpdate, u, e, c, b)
	u.Fire(EVNT_BuffUpdate, u, e, c, b)
	e.Fire(EVNT_BuffUpdate, u, e, c, b)

	return b.Duration > 0
}
