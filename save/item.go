package save

import (
	"errors"
	"time"
)

type Item struct {
	Spell
	Tradable bool
	Stack    uint
	Value    uint
}

func MakeItem(id uint) *Item {
	s := MakeSpell(id)
	i := &Item{
		Spell: *s,
	}

	return i
}

func GetItem(id uint) (*Item, error) {
	if i, ok := Items[id]; i != nil {
		return i, nil
	} else if ok {
		return nil, nil
	}

	i := MakeItem(id)
	err := i.Query()

	if err != nil {
		i = nil
	}

	Items[id] = i
	return i, err
}

func (i *Item) Query() error {
	row := Connection.QueryRow(
		"SELECT script, casttime, cooldown, tradable, stack, value FROM items WHERE id=?",
		i.Id,
	)

	var casttimebuff, cooldownbuff, tradeablebuff int64
	if err := row.Scan(&i.ScriptId, &casttimebuff, &cooldownbuff, &tradeablebuff, &i.Stack, &i.Value); err != nil {
		return err
	}

	if t := time.Duration(casttimebuff); casttimebuff > 0 {
		i.CastTime = t * time.Millisecond
	}
	if t := time.Duration(cooldownbuff); cooldownbuff > 0 {
		i.Cooldown = t * time.Millisecond
	}
	if tradeablebuff > 0 {
		i.Tradable = true
	}

	rows, err := Connection.Query(
		"SELECT key, value FROM animations_items WHERE item=?",
		i.Id,
	)
	if err != nil {
		return err
	}

	for rows.Next() {
		var keybuff, valuebuff string
		err = rows.Scan(&keybuff, &valuebuff)
		if err != nil {
			return err
		}
		i.Animations[keybuff] = valuebuff
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

	return nil
}

func (i *Item) Insert() error {
	return errors.New("save.Item.Insert")
}

func (i *Item) Update() error {
	if err := i.Delete(); err != nil {
		return err
	} else if err := i.Insert(); err != nil {
		return err
	}

	return nil
}

func (i *Item) Delete() error {
	return errors.New("save.Item.Delete")
}

var Items = make(map[uint]*Item)
