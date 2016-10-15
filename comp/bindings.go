package comp

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/util"
	"time"
)

const BINDINGS = "Bindings"

type Bindings map[string]*save.Binding

func init() {
	util.Event.On(save.EVNT_PlayersGet, func(args ...interface{}) {
		bindingsplayerget(args[0].(*save.Player))
	})
	util.Event.On(save.EVNT_PlayersInsert, func(args ...interface{}) {
		bindingsplayerinsert(args[0].(*save.Player))
	})
	util.Event.On(save.EVNT_PlayersDelete, func(args ...interface{}) {
		bindingsplayerdelete(args[0].(string))
	})
}

func (b Bindings) Id() string {
	return BINDINGS
}

func (b Bindings) Update(u *engine.Universe, e *engine.Entity, d time.Duration) {
	for _, binding := range b {
		binding.Update(u, e, d)
	}
}

func (b Bindings) Snapshot() map[string]interface{} {
	return nil
}

func bindingsplayerget(p *save.Player) {
	bindings, err := save.BindingsGetPlayer(p.Username)

	if err != nil {
		logrus.Error(err.Error())
		return
	}

	comp := Bindings(bindings)
	p.Entity.AddComponent(comp)
}

func bindingsplayerinsert(p *save.Player) {
	comp := p.Entity.Component(BINDINGS).(Bindings)

	for _, binding := range comp {
		if err := save.BindingsInsertPlayer(p.Username, binding); err != nil {
			logrus.Error(err.Error())
		}
	}
}

func bindingsplayerdelete(username string) {
	save.BindingsDeletePlayer(username)
}
