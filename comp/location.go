package comp

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/space"
	"github.com/ospokemon/ospokemon/util"
	"time"
)

const LOCATION = "Location"

type Location save.Location

func init() {
	util.Event.On(save.EVNT_PlayersNew, func(args ...interface{}) {
		locationsplayernew(args[0].(*save.Player))
	})
	util.Event.On(save.EVNT_PlayersGet, func(args ...interface{}) {
		locationsplayerget(args[0].(*save.Player))
	})
	util.Event.On(save.EVNT_PlayersInsert, func(args ...interface{}) {
		locationsplayerinsert(args[0].(*save.Player))
	})
}

func (l *Location) Id() string {
	return LOCATION
}

func (l *Location) Snapshot() map[string]interface{} {
	return map[string]interface{}{
		"universe": l.UniverseId,
		"shape":    l.Shape.Snapshot(),
	}
}

func (l *Location) Update(u *engine.Universe, e *engine.Entity, d time.Duration) {
	// TODO
}

func locationsplayernew(p *save.Player) {
	shape := space.Rect{}
	location := save.NewLocation(shape)
	comp := Location(*location)
	p.Entity.AddComponent(&comp)
}

func locationsplayerget(p *save.Player) {
	location, err := save.LocationsGetPlayer(p.Username)

	if err != nil {
		logrus.Error(err.Error())
		return
	}

	comp := Location(*location)
	p.Entity.AddComponent(&comp)
}

func locationsplayerinsert(p *save.Player) {
	comp := p.Entity.Component(LOCATION).(*Location)
	location := save.Location(*comp)

	if err := save.LocationsInsertPlayer(p.Username, &location); err != nil {
		logrus.Error(err.Error())
	}
}
