package cmd

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/comp"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/server"
	"github.com/ospokemon/ospokemon/util"
	"strconv"
	"time"
)

func init() {
	util.Event.On(engine.EVNT_UniverseUpdate, UniverseUpdate)
}

func UniverseUpdate(args ...interface{}) {
	u := args[0].(*engine.Universe)
	d := args[1].(time.Duration)

	mdata := make(map[string]interface{})

	for entityId, entity := range u.Entities {
		entity.Update(u, d)
		mdata[strconv.Itoa(int(entityId))] = entity.Snapshot()
	}

	m := map[string]interface{}{
		"event": "Universe.Update",
		"data":  mdata,
	}

	for _, entity := range u.Entities {
		p := entity.Component(comp.PLAYER).(*comp.Player)

		if p == nil {
			continue
		}

		a := save.Accounts[p.Username]
		s := server.Sessions[a.SessionId]

		snapshot, err := json.Marshal(m)

		if err != nil {
			logrus.Error(err)
		}

		s.Send(string(snapshot))
	}
}
