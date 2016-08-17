package engine

import (
	"github.com/ospokemon/ospokemon/util"
	"time"
)

const EVNT_ActionCast = "ospokemon/engine/Action.Cast"

type Action struct {
	Name     string
	Image    string
	ScriptId string
	CastTime time.Duration
	CoolDown time.Duration
	Timer    *time.Duration
}

func (a *Action) Update(u *Universe, e *Entity, d time.Duration) {
	if a.Timer == nil {
		return
	}

	if *a.Timer < a.CastTime && *a.Timer+d >= a.CastTime {
		a.Cast(u, e)
	}

	*a.Timer = *a.Timer + d
}

func (a *Action) Cast(u *Universe, e *Entity) {
	util.Event.Fire(EVNT_ActionCast, u, e, a)
	u.Fire(EVNT_ActionCast, u, e, a)
	e.Fire(EVNT_ActionCast, u, e, a)
}
