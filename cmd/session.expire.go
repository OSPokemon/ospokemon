package cmd

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/server"
	"github.com/ospokemon/ospokemon/util"
)

func init() {
	util.Event.On(server.EVNT_SessionExpire, SessionExpire)
}

func SessionExpire(args ...interface{}) {
	s := args[0].(*server.Session)
	log := logrus.WithFields(logrus.Fields{
		"SessionId": s.SessionId,
		"Username":  s.Username,
	})

	if err := sessionexpire(s); err != nil {
		log.Error("server.SessionExpire: " + err.Error())
	} else {
		log.Warn("server.SessionExpire")
	}
}

func sessionexpire(s *server.Session) error {
	if s.SessionId < 1 {
		return errors.New("Session already expired")
	}

	delete(server.Sessions, s.SessionId)
	s.SessionId = 0

	if s.Websocket != nil {
		if err := websocketdisconnect(s); err != nil {
			return err
		}
	}
	if a := save.Accounts[s.Username]; a != nil {
		if err := accountlogout(a); err != nil {
			return err
		}
	}

	return nil
}
