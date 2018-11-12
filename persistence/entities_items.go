package persistence

import (
	"github.com/pkg/errors"
	"ospokemon.com"
	"ztaylor.me/log"
)

type tableEntitiesItems struct {
	entity   uint
	universe uint
	item     uint
	amount   int
}

func EntitiesItemsSelect(universe *ospokemon.Universe) (map[uint]*ospokemon.Itemslot, error) {
	rows, err := Connection.Query(
		"SELECT entity, item, amount FROM entities_items WHERE universe=?",
		universe.Id,
	)

	if err != nil {
		return nil, errors.Wrap(err, "entities_items.select")
	}

	entities_items_buff := make([]*tableEntitiesItems, 0)

	for rows.Next() {
		buf := &tableEntitiesItems{}
		err = rows.Scan(&buf.entity, &buf.item, &buf.amount)
		if err != nil {
			return nil, errors.Wrap(err, "entities_items.scan")
		}
		entities_items_buff = append(entities_items_buff, buf)
	}
	rows.Close()

	itemslots := make(map[uint]*ospokemon.Itemslot)

	for _, buf := range entities_items_buff {
		item, err := ospokemon.GetItem(buf.item)
		if err != nil {
			return nil, errors.Wrap(err, "entities_items.select")
		}

		itemslot := ospokemon.BuildItemslot(item, buf.amount)
		itemslots[buf.entity] = itemslot
	}

	log.Add("Universe", universe.Id).Add("Itemslots", entities_items_buff).Debug("entities_items select")

	return itemslots, nil
}
