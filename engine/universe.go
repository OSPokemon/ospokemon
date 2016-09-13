package engine

import (
	"github.com/cznic/mathutil"
	"github.com/ospokemon/ospokemon/util"
	"time"
)

const EVNT_UniverseLoad = "engine/Universe.Load"
const EVNT_UniverseAdd = "engine/Universe.Add"
const EVNT_UniverseRemove = "engine/Universe.Remove"
const EVNT_UniverseUpdate = "engine/Universe.Update"

type Universe struct {
	Id uint
	*Space
	Entities map[uint]Entity
	util.Eventer
	// internals
	bodyIdGen *mathutil.FC32
}

func (u *Universe) GenerateId() uint {
	if u.bodyIdGen == nil {
		bodyIdGen, _ := mathutil.NewFC32(0, 999999, true)
		u.bodyIdGen = bodyIdGen
	}

	return uint(u.bodyIdGen.Next())
}

func (u *Universe) Update(d time.Duration) {
	for _, e := range u.Entities {
		e.Update(u, d)
	}

	util.Event.Fire(EVNT_UniverseUpdate, u)
}

var Multiverse = make(map[uint]*Universe)
