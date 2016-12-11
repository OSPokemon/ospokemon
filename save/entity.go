package save

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"time"
)

type Entities map[uint]*Entity

type Entity struct {
	Id    uint
	Image string
	Components
}

func init() {
	event.On(event.UniverseQuery, func(args ...interface{}) {
		u := args[0].(*Universe)
		entities := make(Entities)

		if err := entities.QueryUniverse(u.Id); err != nil {
			logrus.WithFields(logrus.Fields{
				"Universe": u.Id,
			}).Error("save.Entity: " + err.Error())
		}

		for _, entity := range entities {
			u.Add(entity)
		}
	})
}

func MakeEntity() *Entity {
	return &Entity{
		Components: make(Components),
	}
}

func (e *Entity) Update(u *Universe, d time.Duration) {
	for _, c := range e.Components {
		// c.Update(u, e, d)
	}
}

func (e *Entity) Snapshot() map[string]interface{} {
	c := make(map[string]interface{})
	m := map[string]interface{}{
		"id":    e.Id,
		"image": e.Image,
		"comp":  c,
	}

	for key, comp := range e.Components {
		if compsnap := comp.Snapshot(); compsnap != nil {
			c[key] = compsnap
		}
	}

	return m
}

func (entites Entities) QueryUniverse(universe uint) error {
	rows, err := Connection.Query(
		"SELECT id, image FROM entities_universes WHERE universe=?",
		universe,
	)

	if err != nil {
		return err
	}

	for rows.Next() {
		e := MakeEntity()

		err = rows.Scan(&e.Id, &e.Image)
		if err != nil {
			return err
		}

		entites[e.Id] = e
		e.QueryUniverse(universe)
	}
	rows.Close()

	return nil
}

func (e *Entity) QueryUniverse(universeId uint) {
	event.Fire(event.EntityQuery, e, universeId)
	e.Id = 0 // held temp id, that is deleted now
}
