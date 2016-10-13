package comp

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/util"
	"time"
)

const ACTIONS = "Actions"

type Actions map[uint]*save.Action

func init() {
	util.Event.On(save.EVNT_PlayersGet, func(args ...interface{}) {
		actionsplayerget(args[0].(*save.Player))
	})
	util.Event.On(save.EVNT_PlayersInsert, func(args ...interface{}) {
		actionsplayerinsert(args[0].(*save.Player))
	})
}

func (a Actions) Id() string {
	return ACTIONS
}

func (a Actions) Update(u *engine.Universe, e *engine.Entity, d time.Duration) {
	for _, action := range a {
		action.Update(u, e, d)
	}
}

func (a Actions) Snapshot() map[string]interface{} {
	return nil
}

func actionsplayerget(p *save.Player) {
	actions, err := save.ActionsGetPlayer(p.Username)

	if err != nil {
		logrus.Error(err.Error())
		return
	}

	for spellId := range actions {
		if save.Spells[spellId] == nil {
			if s, err := save.SpellsGet(spellId); err == nil {
				save.Spells[spellId] = s
			} else {
				logrus.Error(err.Error())
			}
		}
	}

	comp := Actions(actions)
	p.Entity.AddComponent(comp)
}

func actionsplayerinsert(p *save.Player) {
	comp := p.Entity.Component(ACTIONS).(Actions)

	for _, action := range comp {
		if err := save.ActionsInsertPlayer(p.Username, action); err != nil {
			logrus.Error(err.Error())
		}
	}
}
