package save

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/part"
	"strconv"
	"time"
)

type Itembag struct {
	Timers map[uint]*time.Duration
	Slots  []*Itemslot
}

func init() {
	event.On(event.PlayerMake, func(args ...interface{}) {
		p := args[0].(*Player)
		bag := MakeItembag(p.BagSize)

		p.AddPart(bag)
	})

	event.On(event.PlayerQuery, func(args ...interface{}) {
		p := args[0].(*Player)
		bag := p.Parts[part.ITEMBAG].(*Itembag)
		err := bag.QueryPlayer(p.Username)

		if err != nil {
			logrus.Error(err.Error())
		}
	})

	event.On(event.PlayerInsert, func(args ...interface{}) {
		p := args[0].(*Player)
		bag := p.Parts[part.ITEMBAG].(*Itembag)
		err := bag.InsertPlayer(p.Username)

		if err != nil {
			logrus.Error(err.Error())
		}
	})

	event.On(event.PlayerDelete, func(args ...interface{}) {
		p := args[0].(*Player)
		bag := p.Parts[part.ITEMBAG].(*Itembag)
		err := bag.DeletePlayer(p.Username)

		if err != nil {
			logrus.Error(err.Error())
		}
	})

	event.On(event.BindingDown, func(args ...interface{}) {
		p := args[0].(*Player)
		binding := args[1].(*Binding)

		if itemslot, ok := binding.Parts[part.ITEMSLOT].(*Itemslot); ok {
			Itembag := p.Parts[part.ITEMBAG].(*Itembag)
			Itembag.Cast(itemslot)
		}
	})
}

func MakeItembag(size uint) *Itembag {
	bag := &Itembag{
		Timers: make(map[uint]*time.Duration),
		Slots:  make([]*Itemslot, size),
	}

	return bag
}

func (b *Itembag) clear() {
	for key := range b.Timers {
		delete(b.Timers, key)
	}
	for pos := range b.Slots {
		b.Slots[pos] = nil
	}
}

func (b *Itembag) Add(slot *Itemslot) bool {
	item, err := GetItem(slot.Item)
	if item == nil {
		if err != nil {
			logrus.Error(err.Error())
		}
		return false
	}

	for _, s := range b.Slots {
		if s == nil {
			continue
		}
		if s.Item != item.Id {
			continue
		}

		if item.Stack < s.Amount+slot.Amount {
			continue
		}

		s.Amount += slot.Amount
		slot.Item = 0
		slot.Amount = 0
		return true
	}

	for id, s := range b.Slots {
		if s == nil {
			slot.Id = id
			b.Slots[id] = slot
			delete(slot.Parts, part.ENTITY)
			return true
		}
	}

	return false
}

func (b *Itembag) Part() string {
	return part.ITEMBAG
}

func (b *Itembag) Cast(slot *Itemslot) {
	if b.Timers[slot.Item] != nil {
		return
	} else if item, err := GetItem(slot.Item); item != nil {
		timer := item.CastTime + item.Cooldown
		b.Timers[slot.Item] = &timer
	} else if err != nil {
		logrus.Error(err.Error())
	}
}

func (b *Itembag) Json(expand bool) (string, map[string]interface{}) {
	data := make(map[string]interface{})

	if expand {
		for pos, slot := range b.Slots {
			key := strconv.Itoa(pos)
			if slot == nil {
				data[key] = nil
				continue
			}

			_, slotData := slot.Json(true)

			if b.Timers[slot.Item] != nil {
				slotData["timer"] = *b.Timers[slot.Item]
			} else {
				slotData["timer"] = 0
			}

			data[key] = slotData
		}
	}

	return "itembag", data
}

func (itembag *Itembag) QueryPlayer(username string) error {
	itembag.clear()

	rows, err := Connection.Query(
		"SELECT pos, item, amount FROM itemslots_players WHERE username=?",
		username,
	)
	if err != nil {
		return err
	}

	for rows.Next() {
		s := MakeItemslot()

		if err = rows.Scan(&s.Id, &s.Item, &s.Amount); err != nil {
			return err
		}

		itembag.Slots[s.Id] = s
	}
	rows.Close()

	rows, err = Connection.Query(
		"SELECT itemid, timer FROM itembags_players WHERE username=?",
		username,
	)
	if err != nil {
		return err
	}

	for rows.Next() {
		var itemidbuff uint
		var timerbuff uint64

		err = rows.Scan(&itemidbuff, &timerbuff)
		if err != nil {
			return err
		}

		if t := time.Duration(timerbuff); timerbuff > 0 {
			itembag.Timers[itemidbuff] = &t
		} else {
			itembag.Timers[itemidbuff] = nil
		}
	}
	rows.Close()

	event.Fire(event.ItembagPlayerQuery, username, itembag)

	return nil
}

func (b *Itembag) InsertPlayer(username string) error {
	for pos, s := range b.Slots {
		if s == nil {
			continue
		}

		_, err := Connection.Exec(
			"INSERT INTO itemslots_players (username, pos, item, amount) VALUES (?, ?, ?, ?)",
			username,
			pos,
			s.Item,
			s.Amount,
		)

		if err != nil {
			return err
		}
	}

	for itemid, timer := range b.Timers {
		timerbuff := 0
		if timer != nil {
			timerbuff = int(*timer)
		}

		_, err := Connection.Exec(
			"INSERT INTO itembags_players (username, itemid, timer) VALUES (?, ?, ?)",
			username,
			itemid,
			timerbuff,
		)

		if err != nil {
			return err
		}
	}

	event.Fire(event.ItembagPlayerInsert, username, b)

	return nil
}

func (b *Itembag) UpdatePlayer(username string) error {
	if err := b.DeletePlayer(username); err != nil {
		return err
	} else if err := b.InsertPlayer(username); err != nil {
		return err
	}

	return nil
}

func (b *Itembag) DeletePlayer(username string) error {
	_, err := Connection.Exec("DELETE FROM itemslots_players WHERE username=?", username)

	if err != nil {
		return err
	}

	_, err = Connection.Exec("DELETE FROM itembags_players WHERE username=?", username)

	if err != nil {
		return err
	}

	event.Fire(event.ItembagPlayerDelete, username)

	return err
}
