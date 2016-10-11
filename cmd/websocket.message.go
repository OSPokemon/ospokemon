package cmd

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/engine"
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
	a := p.Entity.Component(engine.COMP_Actions).(engine.Actions)
	b := p.Entity.Component(engine.COMP_Bindings).(engine.Bindings)

	data := make(map[string]interface{})
	json.Unmarshal([]byte(message), &data)

	action := a[uint(data["spellid"].(float64))]

	binding := &engine.Binding{
		Action: action,
		Timer:  nil,
	}

	b[data["key"].(string)] = binding
}
