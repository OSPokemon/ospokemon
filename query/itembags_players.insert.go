package query

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
)

func ItembagsPlayersInsert(player *game.Player, itembag *game.Itembag) error {
	for pos, itemslot := range itembag.Slots {
		if itemslot == nil {
			continue
		}

		_, err := Connection.Exec(
			"INSERT INTO itemslots_players (username, pos, item, amount) VALUES (?, ?, ?, ?)",
			player.Username,
			pos,
			itemslot.Item.Id,
			itemslot.Amount,
		)

		if err != nil {
			return err
		}
	}

	for itemid, timer := range itembag.Timers {
		timerbuff := 0
		if timer != nil {
			timerbuff = int(*timer)
		}

		_, err := Connection.Exec(
			"INSERT INTO itembags_players (username, itemid, timer) VALUES (?, ?, ?)",
			player.Username,
			itemid,
			timerbuff,
		)

		if err != nil {
			return err
		}
	}

	log.Add("Username", player.Username).Add("Itembag", itembag.Slots).Debug("itembags_players insert")

	event.Fire(event.ItembagsPlayersInsert, player, itembag)

	return nil
}
