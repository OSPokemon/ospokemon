package engine

import (
	"github.com/ospokemon/ospokemon/util"
	"time"
)

const EVNT_EngineUpdate = "ospokemon/engine/Update"

func Run(d time.Duration) {
	for _, universe := range Multiverse {
		universe.Update(d)
	}

	util.Event.Fire(EVNT_EngineUpdate)
}
