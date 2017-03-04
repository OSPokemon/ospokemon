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

	itemslot := game.MakeItemslot()
	err := row.Scan(&itemslot.Item, &itemslot.Amount)

	if err == nil {
		logrus.WithFields(logrus.Fields{
			"Universe": universe.Id,
			"Entity":   entity.Id,
			"Item":     itemslot,
		}).Debug("entities_items select")
	}

	return itemslot, err
}
