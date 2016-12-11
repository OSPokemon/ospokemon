package server

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/run"
	"github.com/ospokemon/ospokemon/save"
)

func ReceiveMessage(s *Session, m *WebsocketMessage) {
	p := save.Players[s.Username]

	s.Refresh()

	if m.Event == "Ping" {
		s.Send("Pong")
	} else if m.Event == "Key.Down" {
		keydown(p, m.Message)
	} else if m.Event == "Key.Up" {
		keyup(p, m.Message)
	} else if m.Event == "Binding.Set" {
		bindingset(p, m.Message)
	} else if m.Event == "Menu.Toggle" {
		menutoggle(p, m.Message)
	} else {
		logrus.WithFields(logrus.Fields{
			"Message":  m,
			"Username": p.Username,
		}).Warn("cmd.WebsocketMessage: Unrecognized message")
	}
}

func keydown(p *save.Player, key string) {
	bindings := p.Entity.Component(save.COMP_Bindings).(save.Bindings)
	binding := bindings[key]

	if binding == nil {
		return
	}

	event.Fire(event.BindingDown, p, binding)
}

func keyup(p *save.Player, key string) {
	bindings := p.Entity.Component(save.COMP_Bindings).(save.Bindings)
	binding := bindings[key]

	if binding == nil {
		return
	}

	event.Fire(event.BindingUp, p, binding)
}

func bindingset(p *save.Player, m string) {
	data := make(map[string]interface{})
	json.Unmarshal([]byte(m), &data)

	bindings := p.Entity.Component(save.COMP_Bindings).(save.Bindings)

	binding := save.MakeBinding(data["key"].(string))

	if spellid, ok := data["spellid"]; ok {
		binding.SpellId = uint(spellid.(float64))

		// binding image

	} else if bagslot, ok := data["bagslot"]; ok {
		binding.BagSlot = int(bagslot.(float64))

		// binding image

	} else if systemid, ok := data["systemid"]; ok {
		binding.SystemId = systemid.(string)

		// binding image

	} else {
		logrus.WithFields(logrus.Fields{
			"Message":  m,
			"Username": p.Username,
		}).Warn("cmd.WebsocketMessage: Unrecognized binding.set message")
	}

	bindings[binding.Key] = binding
}

func menutoggle(p *save.Player, m string) {
	p.Entity.Component(run.COMP_Menus).(*run.Menus).Toggle(m)
}
