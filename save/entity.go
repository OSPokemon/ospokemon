package save

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/part"
	"github.com/ospokemon/ospokemon/space"
	"time"
)

type Entities map[uint]*Entity

type Entity struct {
	Id         uint
	UniverseId uint
	space.Shape
	part.Parts
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

	event.On(event.PlayerMake, func(args ...interface{}) {
		p := args[0].(*Player)
		entity := MakeEntity(&space.Rect{}, p.Parts)
		p.AddPart(entity)
	})

	event.On(event.PlayerQuery, func(args ...interface{}) {
		p := args[0].(*Player)
		entity := p.Parts[part.ENTITY].(*Entity)
		entity.QueryPlayer(p.Username)

		if class, _ := GetClass(p.Class); class != nil {
			r := entity.Shape.(*space.Rect)
			r.Dimension.DX = Classes[p.Class].Dimension.DX
			r.Dimension.DY = Classes[p.Class].Dimension.DY
		}
	})

	event.On(event.PlayerInsert, func(args ...interface{}) {
		p := args[0].(*Player)

		entity := p.Parts[part.ENTITY].(*Entity)
		entity.InsertPlayer(p.Username)
	})
}

func MakeEntity(shape space.Shape, parts part.Parts) *Entity {
	e := &Entity{
		Shape: shape,
		Parts: parts,
	}
	parts.AddPart(e)
	return e
}

func (e *Entity) Part() string {
	return part.ENTITY
}

func (e *Entity) Update(u *Universe, d time.Duration) {
	for _, part := range e.Parts {
		if updater, ok := part.(Updater); ok {
			updater.Update(u, e, d)
		}
	}
}

func (e *Entity) Json(expand bool) (string, map[string]interface{}) {
	data := map[string]interface{}{
		"id":    e.Id,
		"shape": e.Shape.Snapshot(),
	}

	if expand {
		for _, part := range e.Parts {
			if jsoner, ok := part.(Jsoner); ok {
				key, partData := jsoner.Json(false)
				data[key] = partData
			}
		}
	}

	return "entity", data
}

func (entites Entities) QueryUniverse(universe uint) error {
	rows, err := Connection.Query(
		"SELECT id, universe, x, y, dx, dy FROM entities_universes WHERE universe=?",
		universe,
	)

	if err != nil {
		return err
	}

	for rows.Next() {
		r := &space.Rect{}
		e := MakeEntity(r, make(part.Parts))

		err = rows.Scan(&e.Id, &e.UniverseId, &r.Anchor.X, &r.Anchor.Y, &r.Dimension.DX, &r.Dimension.DY)
		if err != nil {
			return err
		}

		entites[e.Id] = e
		event.Fire(event.EntityQuery, e, universe)
		e.Id = 0 // delete temp id
	}
	rows.Close()

	return nil
}

func (e *Entity) QueryPlayer(username string) error {
	r := e.Shape.(*space.Rect)

	row := Connection.QueryRow(
		"SELECT universe, x, y FROM entities_players WHERE username=?",
		username,
	)

	if err := row.Scan(&e.UniverseId, &r.Anchor.X, &r.Anchor.Y); err != nil {
		return err
	}

	return nil
}

func (e *Entity) InsertPlayer(username string) error {
	r := e.Shape.(*space.Rect)

	_, err := Connection.Exec(
		"INSERT INTO entities_players (username, universe, x, y) VALUES (?, ?, ?, ?, ?, ?)",
		username,
		// l.UniverseId,
		r.Anchor.X,
		r.Anchor.Y,
	)

	if err != nil {
		return err
	}

	return nil
}

func (e *Entity) DeletePlayer(username string) error {
	_, err := Connection.Exec("DELETE FROM entities_players WHERE username=?", username)

	return err
}
