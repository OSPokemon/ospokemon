package persistence

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"time"
)

func ItembagsPlayersSelect(player *ospokemon.Player) (*ospokemon.Itembag, error) {
	rows, err := Connection.Query(
		"SELECT sort, item, amount FROM itemslots_players WHERE username=?",
		player.Username,
	)
	if err != nil {
		return nil, err
	}

	itembag := ospokemon.MakeItembag()

	for rows.Next() {
		var amountbuff, sortbuff int
		var itembuff uint

		if err = rows.Scan(&sortbuff, &itembuff, &amountbuff); err != nil {
			return itembag, err
		}

		item, err := ospokemon.GetItem(itembuff)
		if err != nil {
			return itembag, err
		}

		itembag.Slots[itembuff] = ospokemon.BuildItemslot(item, amountbuff)
		itembag.Slots[itembuff].Sort = sortbuff
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

func ItembagsPlayersInsert(player *ospokemon.Player, itembag *ospokemon.Itembag) error {
	for _, itemslot := range itembag.Slots {
		if itemslot == nil {
			continue
		}

		_, err := Connection.Exec(
			"INSERT INTO itemslots_players (username, sort, item, amount) VALUES (?, ?, ?, ?)",
			player.Username,
			itemslot.Sort,
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

func ItembagsPlayersDelete(player *ospokemon.Player) error {
	_, err := Connection.Exec(
		"DELETE FROM itemslots_players WHERE username=?",
		player.Username,
	)

	if err != nil {
		return err
	}

	_, err = Connection.Exec(
		"DELETE FROM itembags_players WHERE username=?",
		player.Username,
	)

	if err != nil {
		log.Add("Username", player.Username).Debug("itembags_players delete")

		event.Fire(event.ItembagsPlayersDelete, player)
	}

	return err
}
