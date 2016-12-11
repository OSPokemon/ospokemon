package save

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"time"
)

const COMP_ItemSlot = "item"

type ItemSlot struct {
	Pos    int
	Item   uint
	Amount uint
}

func init() {
	event.On(event.EntityQuery, func(args ...interface{}) {
		e := args[0].(*Entity)
		universeId := args[1].(uint)

		comp := &ItemSlot{}
		if err := comp.QueryEntityUniverse(e.Id, universeId); err != nil {
			logrus.WithFields(logrus.Fields{
				"Entity":   e.Id,
				"Universe": universeId,
			}).Error("save.ItemSlot(entity): " + err.Error())
			return
		}

		e.AddComponent(comp)
		if item, err := GetItem(comp.Item); item != nil {
			e.Image = item.Image
		} else if err != nil {
			logrus.WithFields(logrus.Fields{
				"Entity":   e.Id,
				"Universe": universeId,
				"Item":     comp.Item,
			}).Error("save.ItemSlot(item): " + err.Error())
		}
	})
}

func (i *ItemSlot) Id() string {
	return COMP_ItemSlot
}

func (i *ItemSlot) Update(u *Universe, e *Entity, d time.Duration) {
}

func (i *ItemSlot) Snapshot() map[string]interface{} {
	data := make(map[string]interface{})

	if item, _ := GetItem(i.Item); item != nil {
		data["pos"] = i.Pos
		data["item"] = item.Snapshot()
		data["amount"] = i.Amount
	}

	return data
}

func (i *ItemSlot) QueryEntityUniverse(entityId uint, universeId uint) error {
	row := Connection.QueryRow(
		"SELECT item, amount FROM entities_items WHERE entity=? AND universe=?",
		entityId,
		universeId,
	)

	if err := row.Scan(&i.Item, &i.Amount); err != nil {
		return err
	}

	return nil
}
