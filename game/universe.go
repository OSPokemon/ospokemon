package game

import (
	"github.com/Sirupsen/logrus"
	"github.com/cznic/mathutil"
	"github.com/ospokemon/ospokemon/space"
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

func (u *Universe) Add(e *Entity) {
	e.Id = u.GenerateId()
	u.Entities[e.Id] = e

	logrus.WithFields(logrus.Fields{
		"Universe": u.Id,
		"Entity":   e.Id,
	}).Debug("game.Universe.Add")

	// event.Fire(event.UniverseAdd, u, e)
}

func (u *Universe) Remove(e *Entity) {
	logrus.WithFields(logrus.Fields{
		"Universe": u.Id,
		"Entity":   e.Id,
	}).Debug("game.Universe.Remove")

	delete(u.Entities, e.Id)
	e.Id = 0

	// event.Fire(event.UniverseRemove, u, e)
}

var Multiverse = make(map[uint]*Universe)