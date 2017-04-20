package ospokemon

import (
	"github.com/cznic/mathutil"
	"ospokemon.com/json"
	"ospokemon.com/log"
	"ospokemon.com/space"
	"time"
)

type Universe struct {
	Id uint
	*Space
	Entities map[uint]*Entity
	Spawners []*Spawner
	Private  bool
	// internals
	bodyIdGen *mathutil.FC32
	Frame     json.Json
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
		Frame:     nil,
	}
}

func (u *Universe) GenerateId() uint {
	return uint(u.bodyIdGen.Next())
}

func (u *Universe) Update(d time.Duration) {
	frame := json.Json{}
	for entityId, entity := range u.Entities {
		if entity == nil {
			continue
		}

		entity.Update(u, d)
		frame[json.StringUint(entityId)] = entity.Json()
	}
	u.Frame = frame

	for _, spawner := range u.Spawners {
		spawner.Update(u, d)
	}
}

func (u *Universe) Add(entities ...*Entity) {
	entitieslog := make([]uint, len(entities))
	for i, e := range entities {
		e.Id = u.GenerateId()
		u.Entities[e.Id] = e
		entitieslog[i] = e.Id
	}
	log.Add("Universe", u.Id).Add("Entity", entitieslog).Debug("ospokemon.Universe.Add")
}

func (u *Universe) AddSpawner(spawner *Spawner) {
	u.Spawners = append(u.Spawners, spawner)
}

func (u *Universe) Remove(e *Entity) {
	log.Add("Universe", u.Id).Add("Entity", e.Id).Debug("ospokemon.Universe.Remove")

	delete(u.Entities, e.Id)
	e.Id = 0
}

var Multiverse = make(map[uint]*Universe)

func GetUniverse(id uint) (*Universe, error) {
	if Multiverse[id] == nil {
		if m, err := Universes.Select(id); err == nil {
			Multiverse[id] = m
		} else {
			return nil, err
		}
	}

	return Multiverse[id], nil
}

var Universes struct {
	Select func(uint) (*Universe, error)
}
