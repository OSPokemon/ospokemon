package save

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/space"
	"time"
)

const COMP_Location = "location"

type Location struct {
	UniverseId uint
	space.Shape
}

func init() {
	event.On(event.PlayerMake, func(args ...interface{}) {
		p := args[0].(*Player)
		location := MakeLocation(&space.Rect{})
		p.Entity.AddComponent(location)
	})
	event.On(event.PlayerQuery, func(args ...interface{}) {
		p := args[0].(*Player)
		location := p.Entity.Component(COMP_Location).(*Location)
		location.QueryPlayer(p.Username)

		if class, _ := GetClass(p.Class); class != nil {
			r := location.Shape.(*space.Rect)
			r.Dimension.DX = Classes[p.Class].Dimension.DX
			r.Dimension.DY = Classes[p.Class].Dimension.DY
		}
	})
	event.On(event.EntityQuery, func(args ...interface{}) {
		e := args[0].(*Entity)
		universeId := args[1].(uint)

		location := MakeLocation(&space.Rect{})
		if err := location.QueryEntityUniverse(e.Id, universeId); err == nil {
			e.AddComponent(location)
		}
	})
	event.On(event.PlayerInsert, func(args ...interface{}) {
		p := args[0].(*Player)

		location := p.Entity.Component(COMP_Location).(*Location)
		location.InsertPlayer(p.Username)
	})
}

func MakeLocation(shape space.Shape) *Location {
	return &Location{
		Shape: shape,
	}
}

func (l *Location) Id() string {
	return COMP_Location
}

func (l *Location) Snapshot() map[string]interface{} {
	return map[string]interface{}{
		"universe": l.UniverseId,
		"shape":    l.Shape.Snapshot(),
	}
}

func (l *Location) Update(u *Universe, e *Entity, d time.Duration) {
	// TODO
}

func (l Location) QueryPlayer(username string) error {
	r := l.Shape.(*space.Rect)

	row := Connection.QueryRow(
		"SELECT universe, x, y FROM locations_players WHERE username=?",
		username,
	)

	if err := row.Scan(&l.UniverseId, &r.Anchor.X, &r.Anchor.Y); err != nil {
		return err
	}

	return nil
}

func (l *Location) QueryEntityUniverse(entityId uint, universeId uint) error {
	r := l.Shape.(*space.Rect)

	row := Connection.QueryRow(
		"SELECT x, y, dx, dy FROM entities_locations WHERE entity=? AND universe=?",
		entityId,
		universeId,
	)

	if err := row.Scan(&r.Anchor.X, &r.Anchor.Y, &r.Dimension.DX, &r.Dimension.DY); err != nil {
		return err
	}

	l.UniverseId = universeId

	return nil
}

func (l Location) InsertPlayer(username string) error {
	r := l.Shape.(*space.Rect)

	_, err := Connection.Exec(
		"INSERT INTO locations_players (username, universe, x, y) VALUES (?, ?, ?, ?, ?, ?)",
		username,
		l.UniverseId,
		r.Anchor.X,
		r.Anchor.Y,
	)

	if err != nil {
		return err
	}

	return nil
}

func (l *Location) UpdatePlayer(username string) error {
	if err := l.DeletePlayer(username); err != nil {
		return err
	} else if err := l.InsertPlayer(username); err != nil {
		return err
	}

	return nil
}

func (l *Location) DeletePlayer(username string) error {
	_, err := Connection.Exec("DELETE FROM locations_players WHERE username=?", username)

	return err
}
