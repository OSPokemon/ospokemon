package persistence

import (
	"ospokemon.com"
	"ztaylor.me/log"
)

func EntitiesItemsSelect(universe *ospokemon.Universe) (map[uint]*ospokemon.Itemslot, error) {
	rows, err := Connection.Query(
		"SELECT entity, item, amount FROM entities_items WHERE universe=?",
		universe.Id,
	)

	if err != nil {
		return nil, err
	}

	itemslotslog := make(map[uint]uint)
	itemslots := make(map[uint]*ospokemon.Itemslot)

	for rows.Next() {
		var entityId, itembuff uint
		var amountbuff int
		err = rows.Scan(&entityId, &itembuff, &amountbuff)
		if err != nil {
			return nil, err
		}

		item, err := ospokemon.GetItem(itembuff)
		if err != nil {
			return nil, err
		}

		itemslot := ospokemon.BuildItemslot(item, amountbuff)
		itemslots[entityId] = itemslot
		itemslotslog[entityId] = itemslot.Item.Id
	}

	log.Add("Universe", universe.Id).Add("Itemslots", itemslotslog).Debug("entities_items select")

	return itemslots, nil
}
