package save

import (
	"github.com/cznic/mathutil"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/space"
	"strconv"
	"time"
)

type Universe struct {
	Id uint
	*Space
	Entities map[uint]*Entity
	Private  bool
	// internals
	bodyIdGen *mathutil.FC32
}

type Space struct {
	Dimension space.Vector
	Division  *space.Line
	Sub       *[2]*Space
	Entities  []Entity
}

func MakeUniverse(universeid uint) *Universe {
	bodyIdGen, _ := mathutil.NewFC32(0, 999999, true)

	return &Universe{
		Id: universeid,
		Space: &Space{
			Dimension: space.Vector{},
			Division:  nil,
			Sub:       nil,
			Entities:  make([]Entity, 0),
		},
		Entities:  make(map[uint]*Entity),
		bodyIdGen: bodyIdGen,
	}
}

func GetUniverse(id uint) (*Universe, error) {
	if u, ok := Multiverse[id]; u != nil {
		return u, nil
	} else if ok {
		return nil, nil
	}

	u := MakeUniverse(id)
	err := u.Query()

	if err != nil {
		u = nil
	}

	Multiverse[id] = u
	return u, err
}

func (u *Universe) GenerateId() uint {
	return uint(u.bodyIdGen.Next())
}

func (u *Universe) Update(d time.Duration) {
	for _, e := range u.Entities {
		if e == nil {
			continue
		}

		e.Update(u, d)
	}
}

func (u *Universe) Snapshot() map[string]interface{} {
	data := make(map[string]interface{})

	for entityId, entity := range u.Entities {
		key := strconv.Itoa(int(entityId))

		if entity == nil {
			data[key] = nil
		} else {
			data[key] = entity.Snapshot()
		}
	}

	return data
}

func (u *Universe) Add(e *Entity) {
	e.Id = u.GenerateId()
	u.Entities[e.Id] = e
}

func (u *Universe) Remove(e *Entity) {
	u.Entities[e.Id] = nil
	e.Id = 0
}

func (u *Universe) Query() error {
	row := Connection.QueryRow(
		"SELECT dx, dy, private FROM universes WHERE id=?",
		u.Id,
	)

	if err := row.Scan(&u.Space.Rect.Dimension.DX, &u.Space.Rect.Dimension.DY, &u.Private); err != nil {
		return err
	}

	event.Fire(event.UniverseQuery, u)

	return nil
}

var Multiverse = make(map[uint]*Universe)
