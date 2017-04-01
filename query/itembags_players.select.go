package query

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"time"
)

func ItembagsPlayersSelect(player *ospokemon.Player) (*ospokemon.Itembag, error) {
	rows, err := Connection.Query(
		"SELECT pos, item, amount FROM itemslots_players WHERE username=?",
		player.Username,
	)
	if err != nil {
		return nil, err
	}

	itembag := ospokemon.MakeItembag(player.BagSize)

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

		itemslot := ospokemon.BuildItemslot(item, amountbuff)
		itemslot.Id = idbuff
		itembag.Slots[idbuff] = itemslot
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
