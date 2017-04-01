package query

import (
	"ospokemon.com/game"
	"ospokemon.com/log"
)

func EntitiesItemsSelect(universe *game.Universe) (map[uint]*game.Itemslot, error) {
	rows, err := Connection.Query(
		"SELECT entity, item, amount FROM entities_items WHERE universe=?",
		universe.Id,
	)

	if err != nil {
		return nil, err
	}

	itemslots := make(map[uint]*game.Itemslot)

	for rows.Next() {
		var entityId, itembuff uint
		var amountbuff int
		err = rows.Scan(&entityId, &itembuff, &amountbuff)
		if err != nil {
			return nil, err
		}

		item, err := GetItem(itembuff)
		if err != nil {
			return nil, err
		}

		itemslot := game.BuildItemslot(item, amountbuff)
		itemslots[entityId] = itemslot
	}

	log.Add("Universe", universe.Id).Add("Itemslots", itemslots).Debug("entities_items select")

	return itemslots, nil
}
