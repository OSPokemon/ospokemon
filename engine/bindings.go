package engine

import (
	"github.com/ospokemon/ospokemon/util"
	"time"
)

const COMP_Bindings = "engine/Bindings"
const EVNT_BindingsBind = "engine/Bindings.Bind"
const EVNT_BindingsUpdate = "engine/Bindings.Update"

type Bindings map[string]*Action

func (b *Bindings) Id() string {
	return COMP_Bindings
}

func (b Bindings) Update(u *Universe, e *Entity, d time.Duration) {
	for _, a := range b {
		a.Update(u, e, d)
	}

	util.Event.Fire(EVNT_BindingsUpdate, u, e, b)
	u.Fire(EVNT_BindingsUpdate, u, e, b)
	e.Fire(EVNT_BindingsUpdate, u, e, b)
}

func (b Bindings) Bind(s string, e *Entity, a *Action) {
	b[s] = a

	util.Event.Fire(EVNT_BindingsBind, e, b, a)
	e.Fire(EVNT_BindingsBind, e, b, a)
}
