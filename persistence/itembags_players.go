package persistence

import (
	"time"

	"github.com/ospokemon/ospokemon"
	"github.com/pkg/errors"
)

type tableItemslotsPlayers struct {
	username string
	item     uint
	amount   int
	sort     int
}

func ItembagsPlayersSelect(player *ospokemon.Player) (*ospokemon.Itembag, error) {
	rows, err := Connection.Query(
		"SELECT sort, item, amount FROM itemslots_players WHERE username=?",
		player.Username,
	)
	if err != nil {
		return nil, errors.Wrap(err, "itemslots_players.select")
	}

	itembagbuf := make([]*tableItemslotsPlayers, 0)

	for rows.Next() {
		var buf tableItemslotsPlayers

		if err = rows.Scan(&buf.sort, &buf.item, &buf.amount); err != nil {
			return nil, errors.Wrap(err, "itemslots_players.scan")
		}
		itembagbuf = append(itembagbuf, &buf)
	}
	rows.Close()

	itembag := ospokemon.MakeItembag()

	for _, buf := range itembagbuf {
		item, err := ospokemon.GetItem(buf.item)
		if err != nil {
			return nil, errors.Wrap(err, "itembags_players.getitem")
		}

		itembag.Slots[buf.item] = ospokemon.BuildItemslot(item, buf.amount)
		itembag.Slots[buf.item].Sort = buf.sort
	}

	rows, err = Connection.Query(
		"SELECT itemid, timer FROM itembags_players WHERE username=?",
		player.Username,
	)
	if err != nil {
		return itembag, errors.Wrap(err, "itembags_players.select")
	}

	for rows.Next() {
		var itembuff uint
		var timebuff uint64

		err = rows.Scan(&itembuff, &timebuff)
		if err != nil {
			return nil, errors.Wrap(err, "itembags_players.scan")
		}

		if t := time.Duration(timebuff); timebuff > 0 {
			itembag.Timers[itembuff].Set(t)
		} else {
			itembag.Timers[itembuff] = nil
		}
	}
	rows.Close()

	ospokemon.LOG().Add("Username", player.Username).Add("Itembag", itembag.GetItems()).Debug("itembags_players select")
	return itembag, nil
}

func ItembagsPlayersInsert(player *ospokemon.Player) error {
	itembag := player.GetItembag()
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

	ospokemon.LOG().Add("Username", player.Username).Add("Itembag", itembag.Slots).Debug("itembags_players insert")
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
		return err
	}

	ospokemon.LOG().Add("Username", player.Username).Debug("itembags_players delete")
	return err
}
