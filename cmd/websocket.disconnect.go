package cmd

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/comp"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/server"
	"github.com/ospokemon/ospokemon/util"
)

func init() {
	util.Event.On(server.EVNT_WebsocketDisconnect, WebsocketDisconnect)
}

func WebsocketDisconnect(args ...interface{}) {
	s := args[0].(*server.Session)
	log := logrus.WithFields(logrus.Fields{
		"SessionId": s.SessionId,
		"Username":  s.Username,
	})

	if err := websocketdisconnect(s); err != nil {
		log.Error("server.WebsocketDisconnect: " + err.Error())
	} else {
		log.Info("server.WebsocketDisconnect")
	}
}

func websocketdisconnect(s *server.Session) error {
	if s.Websocket == nil {
		return errors.New("Websocket already closed")
	}

	s.Websocket.Close()
	s.Websocket = nil

	if p := save.Players[s.Username]; p == nil {
		return errors.New("server.WebsocketDisconnect: Player missing")
	} else {
		l := p.Entity.Component(comp.LOCATION).(*comp.Location)

		if err := universeremove(p.Entity, l); err != nil {
			return err
		}
	}

	return nil
}
