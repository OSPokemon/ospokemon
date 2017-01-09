package save

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/part"
	"time"
)

type Itemslot struct {
	Id     int
	Item   uint
	Amount uint
	Timer  *time.Duration // TODO
	part.Parts
}

func MakeItemslot() *Itemslot {
	itemslot := &Itemslot{
		Id:    -1,
		Parts: make(part.Parts),
	}

	itemslot.AddPart(itemslot)

	event.Fire(event.ItemslotMake, itemslot)

	return itemslot
}

func init() {
	event.On(event.EntityQuery, func(args ...interface{}) {
		e := args[0].(*Entity)
		universeId := args[1].(uint)

		is := MakeItemslot()
		if err := is.QueryEntityUniverse(e.Id, universeId); err != nil {
			logrus.WithFields(logrus.Fields{
				"Entity":   e.Id,
				"Universe": universeId,
			}).Error("save.Itemslot(entity): " + err.Error())
			return
		}

		for _, part := range is.Parts {
			e.AddPart(part)
		}
		e.AddPart(is)
		is.Parts = e.Parts
	})
}

func (i *Itemslot) Part() string {
	return part.ITEMSLOT
}

func (i *Itemslot) Update(u *Universe, e *Entity, d time.Duration) {
}

func (i *Itemslot) Json(expand bool) (string, map[string]interface{}) {
	data := map[string]interface{}{
		"id":     i.Id,
		"amount": i.Amount,
	}

	if item, _ := GetItem(i.Item); item != nil {
		data["item"] = item.Snapshot()
	} else {
		data["item"] = i.Item
	}

	if expand {
		for _, part := range i.Parts {
			if jsoner, ok := part.(Jsoner); ok {
				key, partData := jsoner.Json(false)
				data[key] = partData
			}
		}
	}

	return "itemslot", data
}

func (i *Itemslot) QueryEntityUniverse(entityId uint, universeId uint) error {
	row := Connection.QueryRow(
		"SELECT item, amount FROM entities_items WHERE entity=? AND universe=?",
		entityId,
		universeId,
	)

	if err := row.Scan(&i.Item, &i.Amount); err != nil {
		return err
	}

	event.Fire(event.ItemslotEntityQuery, universeId, entityId, i)

	return nil
}
