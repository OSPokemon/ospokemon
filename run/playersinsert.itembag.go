package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.PlayersInsert, PlayersInsertItembag)
}

func PlayersInsertItembag(args ...interface{}) {
	player := args[0].(*game.Player)
	itembag := player.GetItembag()

	if itembag == nil {
		itembag = game.MakeItembag(player.BagSize)
		log.Add("Username", player.Username).Debug("players insert itembag: grant empty bag")
	}

	err := query.ItembagsPlayersInsert(player, itembag)

	if err != nil {
		log.Add("Player", player.Username).Add("Itembag", itembag.GetItems()).Add("Error", err.Error()).Error("players insert itembag")
	}
}
