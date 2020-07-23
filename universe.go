package ospokemon

import (
	"time"

	"github.com/cznic/mathutil"
	"github.com/ospokemon/ospokemon/space"
	"taylz.io/types"
)

type Universe struct {
	Id uint
	*Space
	Entities map[uint]*Entity
	Spawners []*Spawner
	Private  bool
	// internals
	bodyIdGen *mathutil.FC32
	FullFrame types.Dict
}

type Space struct {
	Dimension space.Vector
	Division  *space.Line
	Sub       *[2]*Space
	Entities  Entities
}

func MakeUniverse(universeid uint) *Universe {
	bodyIdGen, _ := mathutil.NewFC32(0, 999999, true)

	return &Universe{
		Id: universeid,
		Space: &Space{
			Dimension: space.Vector{},
			Division:  nil,
			Sub:       nil,
			Entities:  Entities{},
		},
		Entities:  make(map[uint]*Entity),
		bodyIdGen: bodyIdGen,
		FullFrame: nil,
	}
}

func (u *Universe) GenerateId() uint {
	return uint(u.bodyIdGen.Next())
}

func (u *Universe) Update(d time.Duration) {
	frame := types.Dict{}
	for entityId, entity := range u.Entities {
		if entity == nil {
			continue
		}

		entity.Update(u, d)
		frame[types.StringUint(entityId)] = entity.Json()
	}
	u.FullFrame = frame

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

func GetUniverse(id uint) (u *Universe, err error) {
	if u = Universes.Cache[id]; u == nil {
		if u, err = Universes.Select(id); err == nil {
			Universes.Cache[id] = u
		} else {
			log.With(types.Dict{
				"Universe": id,
				"Error":    err.Error(),
			}).Error("ospokemon.GetUniverse: create")
		}
	}
	return
}

var Universes = struct {
	Cache  map[uint]*Universe
	Select func(uint) (*Universe, error)
}{make(map[uint]*Universe), nil}
