package save

import (
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/util"
)

const EVNT_PlayersNew = "save.PlayersNew"

func PlayersNew(username string) *Player {
	p := &Player{
		Username:   username,
		Level:      0,
		Experience: 0,
		Money:      0,
		Entity:     engine.MakeEntity(),
	}

	util.Event.Fire(EVNT_PlayersNew, p)

	return p
}
