package query

import (
	"ospokemon.com"
	"ospokemon.com/log"
	"time"
)

func ItemsSelect(id uint) (*ospokemon.Item, error) {
	row := Connection.QueryRow(
		"SELECT id, script, casttime, cooldown, tradable, stack, value, dx, dy FROM items WHERE id=?",
		id,
	)

	item := ospokemon.MakeItem()

	var casttimebuff, cooldownbuff, tradeablebuff int64
	if err := row.Scan(&item.Spell.Id, &item.Script, &casttimebuff, &cooldownbuff, &tradeablebuff, &item.Stack, &item.Value, &item.Dimension.DX, &item.Dimension.DY); err != nil {
		return item, err
	}

	if t := time.Duration(casttimebuff); casttimebuff > 0 {
		item.CastTime = t * time.Millisecond
	}
	if t := time.Duration(cooldownbuff); cooldownbuff > 0 {
		item.Cooldown = t * time.Millisecond
	}
	if tradeablebuff > 0 {
		item.Tradable = true
	}

	rows, err := Connection.Query(
		"SELECT key, value FROM animations_items WHERE item=?",
		id,
	)
	if err != nil {
		return item, err
	}

	for rows.Next() {
		var keybuff, valuebuff string
		err = rows.Scan(&keybuff, &valuebuff)
		if err != nil {
			return item, err
		}
		item.Animations[keybuff] = valuebuff
	}
	rows.Close()

	// TODO get item data
	// rows, err = Connection.Query(
	// 	"SELECT key, value FROM items_data WHERE item=?",
	// 	i.Id,
	// )
	// if err != nil {
	// 	return err
	// }

	// for rows.Next() {
	// 	var keybuff, valuebuff string
	// 	err = rows.Scan(&keybuff, &valuebuff)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	i.Data[keybuff] = valuebuff
	// }
	// rows.Close()

	ospokemon.Items[id] = item

	log.Add("Item", id).Info("items select")

	return item, nil
}
