package save

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"strconv"
	"time"
)

const COMP_Bag = "save.Bag"

type Bag struct {
	Timers map[uint]*time.Duration
	Slots  []*ItemSlot
}

func init() {
	event.On(event.PlayerMake, func(args ...interface{}) {
		p := args[0].(*Player)
		bag := MakeBag(p.BagSize)

		p.Entity.AddComponent(bag)
	})

	event.On(event.PlayerQuery, func(args ...interface{}) {
		p := args[0].(*Player)
		bag := p.Entity.Component(COMP_Bag).(*Bag)
		bag.QueryPlayer(p.Username)
	})

	event.On(event.PlayerInsert, func(args ...interface{}) {
		p := args[0].(*Player)
		bag := p.Entity.Component(COMP_Bag).(*Bag)
		bag.InsertPlayer(p.Username)
	})

	event.On(event.PlayerDelete, func(args ...interface{}) {
		p := args[0].(*Player)
		bag := p.Entity.Component(COMP_Bag).(*Bag)
		bag.DeletePlayer(p.Username)
	})

	event.On(event.BindingDown, func(args ...interface{}) {
		p := args[0].(*Player)
		binding := args[1].(*Binding)

		if binding.BagSlot > 0 {
			p.Entity.Component(COMP_Bag).(*Bag).Cast(binding)
		}
	})
}

func MakeBag(size uint) *Bag {
	bag := &Bag{
		Timers: make(map[uint]*time.Duration),
		Slots:  make([]*ItemSlot, size),
	}

	return bag
}

func (b *Bag) clear() {
	for key := range b.Timers {
		delete(b.Timers, key)
	}
	for pos := range b.Slots {
		b.Slots[pos] = nil
	}
}

func (b *Bag) Cast(binding *Binding) {
	slot := b.Slots[binding.BagSlot]
	if slot == nil || b.Timers[slot.Item] != nil {
		return
	} else if item, err := GetItem(slot.Item); item != nil {
		timer := item.CastTime + item.Cooldown
		b.Timers[slot.Item] = &timer
	} else if err != nil {
		logrus.Error(err.Error())
	}
}

func (b *Bag) Add(slot *ItemSlot) bool {
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

	for pos, s := range b.Slots {
		if s == nil {
			addSlot := &ItemSlot{
				Pos:    pos,
				Item:   slot.Item,
				Amount: slot.Amount,
			}
			b.Slots[pos] = addSlot
			slot.Item = 0
			slot.Amount = 0
			return true
		}
	}

	return false
}

func (b *Bag) Id() string {
	return COMP_Bag
}

func (b *Bag) Update(u *Universe, e *Entity, d time.Duration) {
	// TODO
}

func (b *Bag) Snapshot() map[string]interface{} {
	return nil
}

func (b *Bag) SnapshotDetail() map[string]interface{} {
	data := make(map[string]interface{})

	for pos, slot := range b.Slots {
		key := strconv.Itoa(pos)
		if slot == nil {
			data[key] = nil
			continue
		}

		data[key] = slot.Snapshot()
	}

	return data
}

func (b *Bag) QueryPlayer(username string) error {
	b.clear()

	rows, err := Connection.Query(
		"SELECT pos, item, amount FROM itemslots_players WHERE username=?",
		username,
	)
	if err != nil {
		return err
	}

	for rows.Next() {
		s := &ItemSlot{}

		err = rows.Scan(&s.Pos, &s.Item, &s.Amount)
		if err != nil {
			return err
		}

		b.Slots[s.Pos] = s
	}
	rows.Close()

	rows, err = Connection.Query(
		"SELECT itemid, timer FROM bags_players WHERE username=?",
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
			b.Timers[itemidbuff] = &t
		} else {
			b.Timers[itemidbuff] = nil
		}
	}
	rows.Close()

	return nil
}

func (b *Bag) InsertPlayer(username string) error {
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
			"INSERT INTO bags_players (username, itemid, timer) VALUES (?, ?, ?)",
			username,
			itemid,
			timerbuff,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Bag) UpdatePlayer(username string) error {
	if err := b.DeletePlayer(username); err != nil {
		return err
	} else if err := b.InsertPlayer(username); err != nil {
		return err
	}

	return nil
}

func (b *Bag) DeletePlayer(username string) error {
	_, err := Connection.Exec("DELETE FROM itemslots_players WHERE username=?", username)

	if err != nil {
		return err
	}

	_, err = Connection.Exec("DELETE FROM bags_players WHERE username=?", username)

	return err
}
