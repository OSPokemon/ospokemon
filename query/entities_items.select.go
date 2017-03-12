package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/game"
)

func EntitiesItemsSelect(entity *game.Entity, universe *game.Universe) (*game.Itemslot, error) {
	row := Connection.QueryRow(
		"SELECT item, amount FROM entities_items WHERE entity=? AND universe=?",
		entity.Id,
		universe.Id,
	)

	var itembuff uint
	var amountbuff int
	err := row.Scan(&itembuff, &amountbuff)

	item, err := GetItem(itembuff)
	if err != nil {
		return nil, err
	}

	itemslot := game.BuildItemslot(0, item, amountbuff)

	logrus.WithFields(logrus.Fields{
		"Universe": universe.Id,
		"Entity":   entity.Id,
		"Item":     item.Id,
	}).Debug("entities_items select")

	return itemslot, nil
}
