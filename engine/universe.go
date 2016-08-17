package engine

import (
	"github.com/cznic/mathutil"
	"github.com/ospokemon/ospokemon/util"
	"time"
)

const ENVT_UniverseGenerateId = "ospokemon/engine/Universe.GenerateId"

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

	id := uint(u.bodyIdGen.Next())

	util.Event.Fire(ENVT_UniverseGenerateId, u, id)
	u.Fire(ENVT_UniverseGenerateId, u, id)

	return id
}

func (u *Universe) Update(d time.Duration) {
	for _, e := range u.Entities {
		e.Update(u, d)
	}
}

var Multiverse = make(map[uint]*Universe)
