package cmd

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/comp"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/server"
	"github.com/ospokemon/ospokemon/util"
)

func init() {
	util.Event.On(server.EVNT_WebsocketMessage, WebsocketMessage)
}

func WebsocketMessage(args ...interface{}) {
	s := args[0].(*server.Session)
	m := args[1].(server.WebsocketMessage)

	p := save.Players[s.Username]

	s.Refresh()

	if m.Event == "Ping" {
		s.Send("Pong")
	}
	if m.Event == "Binding.Add" {
		bindingadd(p, m.Message)
	}
	if m.Event == "Binding.Cast" {
		bindingcast(p, m.Message)
	}
}

func bindingadd(p *save.Player, message string) {
	b := p.Entity.Component(comp.BINDINGS).(comp.Bindings)

	data := make(map[string]interface{})
	json.Unmarshal([]byte(message), &data)

	binding := &save.Binding{
		Key:     data["key"].(string),
		SpellId: uint(data["spellid"].(float64)),
		Timer:   nil,
	}

	b[binding.Key] = binding
}

func bindingcast(p *save.Player, message string) {
	data := make(map[string]interface{})
	json.Unmarshal([]byte(message), &data)

	bindings := p.Entity.Component(comp.BINDINGS).(comp.Bindings)
	actions := p.Entity.Component(comp.ACTIONS).(comp.Actions)

	key := data["key"].(string)

	if binding := bindings[key]; binding != nil {
		if action := actions[binding.SpellId]; action != nil {
			if action.Timer != nil {
				logrus.WithFields(logrus.Fields{
					"Key":      key,
					"SpellId":  action.SpellId,
					"Username": p.Username,
				}).Warn("cmd.WebsocketMessage: Action fired already running")
				return
			}

			if spell := save.Spells[action.SpellId]; spell != nil {
				timer := spell.CastTime + spell.Cooldown
				action.Timer = &timer
			} else {
				logrus.WithFields(logrus.Fields{
					"Key":      key,
					"SpellId":  action.SpellId,
					"Username": p.Username,
				}).Warn("cmd.WebsocketMessage: Spell lookup failed")
			}
		} else {
			logrus.WithFields(logrus.Fields{
				"Key":      key,
				"SpellId":  binding.SpellId,
				"Username": p.Username,
			}).Warn("cmd.WebsocketMessage: Action lookup failed")
		}
	} else {
		// binding == nil
	}
}
