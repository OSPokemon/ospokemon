package server

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/part"
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
	} else {
		logrus.WithFields(logrus.Fields{
			"Message":  m,
			"Username": p.Username,
		}).Warn("Websocket: Unrecognized message type")
	}
}

func keydown(p *save.Player, key string) {
	bindings := p.Parts[part.BINDINGS].(save.Bindings)
	binding := bindings[key]

	if binding == nil {
		return
	}

	event.Fire(event.BindingDown, p, binding)
}

func keyup(p *save.Player, key string) {
	bindings := p.Parts[part.BINDINGS].(save.Bindings)
	binding := bindings[key]

	if binding == nil {
		return
	}

	event.Fire(event.BindingUp, p, binding)
}

func bindingset(p *save.Player, m string) {
	data := make(map[string]interface{})
	json.Unmarshal([]byte(m), &data)

	key := data["key"].(string)

	bindings := p.Parts[part.BINDINGS].(save.Bindings)

	binding := save.MakeBinding()
	binding.Key = key

	if data["spellid"] != nil {
		spellId := uint(data["spellid"].(float64))
		actions := p.Parts[part.ACTIONS].(save.Actions)

		if action := actions[spellId]; action != nil {
			action.Parts[part.BINDINGS].(save.Bindings)[binding.Key] = binding
			binding.Parts = action.Parts
		}
	} else if data["itemid"] != nil {
		itemid := uint(data["itemid"].(float64))
		itembag := p.Parts[part.ITEMBAG].(*save.Itembag)

		for _, itemslot := range itembag.Slots {
			if itemslot == nil || itemid != itemslot.Item {
				continue
			}

			itemslot.Parts[part.BINDINGS].(save.Bindings)[binding.Key] = binding
			binding.Parts = itemslot.Parts
		}
	} else if systemid, ok := data["systemid"]; ok {
		binding.SystemId = systemid.(string)

		// binding image

	} else {
		logrus.WithFields(logrus.Fields{
			"Message":  m,
			"Username": p.Username,
		}).Warn("Websocket: Unrecognized Binding.Set message")
	}

	bindings[binding.Key] = binding
}
