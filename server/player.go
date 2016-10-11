package server

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/save"
	"net/http"
	"strconv"
)

var PlayerHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	s := readsession(r)

	if r.Method != "GET" || s == nil {
		w.WriteHeader(404)
		return
	}

	p := save.Players[s.Username]

	if p == nil {
		w.WriteHeader(404)
		logrus.WithFields(logrus.Fields{
			"SessionId": s.SessionId,
			"Username":  s.Username,
		}).Warn("server.PlayerHandler: session exists for missing player")
		return
	}

	bdata := make(map[string]interface{})
	bindings := p.Entity.Component(engine.COMP_Bindings).(engine.Bindings)
	for key, binding := range bindings {
		bindingsnap := binding.Snapshot()
		bindingsnap["key"] = key
		bdata[key] = bindingsnap
	}

	adata := make(map[string]interface{})
	actions := p.Entity.Component(engine.COMP_Actions).(engine.Actions)
	for spellid, action := range actions {
		adata[strconv.Itoa(int(spellid))] = action.Snapshot()
	}

	m := p.Snapshot()
	m["bindings"] = bdata
	m["actions"] = adata

	snapshot, err := json.Marshal(m)

	if err != nil {
		logrus.Error(err)
	}

	w.Write(snapshot)
})
