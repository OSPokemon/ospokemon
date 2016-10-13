package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/comp"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/server"
	"github.com/ospokemon/ospokemon/util"
)

func init() {
	util.Event.On(server.EVNT_WebsocketConnect, WebsocketConnect)
}

func WebsocketConnect(args ...interface{}) {
	s := args[0].(*server.Session)
	log := logrus.WithFields(logrus.Fields{
		"SessionId": s.SessionId,
		"Username":  s.Username,
	})

	if err := websocketconnect(s); err != nil {
		log.Error("cmd.WebsocketConnect: " + err.Error())
	} else {
		log.Info("cmd.WebsocketConnect")
	}
}

func websocketconnect(s *server.Session) error {
	if save.Players[s.Username] == nil {
		if p, err := save.PlayersGet(s.Username); err != nil {
			return err
		} else {
			save.Players[p.Username] = p
		}
	}

	p := save.Players[s.Username]
	l := p.Entity.Component(comp.LOCATION).(*comp.Location)

	if err := universeadd(p.Entity, l); err != nil {
		return err
	}

	return nil
}
