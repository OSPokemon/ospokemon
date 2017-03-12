package query

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
	"time"
)

func ItembagsPlayersSelect(player *game.Player) (*game.Itembag, error) {
	rows, err := Connection.Query(
		"SELECT pos, item, amount FROM itemslots_players WHERE username=?",
		player.Username,
	)
	if err != nil {
		return nil, err
	}

	itembag := game.MakeItembag(player.BagSize)

	for rows.Next() {
		var idbuff, amountbuff int
		var itembuff uint

		if err = rows.Scan(&idbuff, &itembuff, &amountbuff); err != nil {
			return itembag, err
		}

		item, err := GetItem(itembuff)
		if err != nil {
			return itembag, err
		}

		itembag.Slots[idbuff] = game.BuildItemslot(idbuff, item, amountbuff)
	}
	rows.Close()

	rows, err = Connection.Query(
		"SELECT itemid, timer FROM itembags_players WHERE username=?",
		player.Username,
	)
	if err != nil {
		return itembag, err
	}

	for rows.Next() {
		var itembuff uint
		var timebuff uint64

		err = rows.Scan(&itembuff, &timebuff)
		if err != nil {
			return itembag, err
		}

		if t := time.Duration(timebuff); timebuff > 0 {
			itembag.Timers[itembuff] = &t
		} else {
			itembag.Timers[itembuff] = nil
		}
	}
	rows.Close()

	log.Add("Username", player.Username).Add("Itembag", itembag.Slots).Debug("itembags_players select")

	event.Fire(event.ItembagsPlayersSelect, player, itembag)

	return itembag, nil
}
