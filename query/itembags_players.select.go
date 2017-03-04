package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
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
		itemslot := game.MakeItemslot()

		if err = rows.Scan(&itemslot.Id, &itemslot.Item, &itemslot.Amount); err != nil {
			return itembag, err
		}

		event.Fire(event.ItemslotBuild, itemslot)

		itembag.Slots[itemslot.Id] = itemslot
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

	logrus.WithFields(logrus.Fields{
		"Username": player.Username,
		"Itembag":  itembag.Slots,
	}).Debug("itembags_players select")

	event.Fire(event.ItembagsPlayersSelect, player, itembag)

	return itembag, nil
}
