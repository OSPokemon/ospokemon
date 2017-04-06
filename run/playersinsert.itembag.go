package run

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"ospokemon.com/persistence"
)

func init() {
	event.On(event.PlayersInsert, PlayersInsertItembag)
}

func PlayersInsertItembag(args ...interface{}) {
	player := args[0].(*ospokemon.Player)
	itembag := player.GetItembag()

	if itembag == nil {
		itembag = ospokemon.MakeItembag()
		log.Add("Username", player.Username).Debug("players insert itembag: grant empty bag")
	}

	err := persistence.ItembagsPlayersInsert(player, itembag)

	if err != nil {
		log.Add("Player", player.Username).Add("Itembag", itembag.GetItems()).Add("Error", err.Error()).Error("players insert itembag")
	}
}
