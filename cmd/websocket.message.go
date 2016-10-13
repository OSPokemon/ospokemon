package cmd

import (
	"encoding/json"
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

	if m.Event == "Binding.Add" {
		bindingadd(p, m.Message)
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
