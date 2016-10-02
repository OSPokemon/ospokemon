package cmd

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/server"
	"github.com/ospokemon/ospokemon/util"
)

func init() {
	util.Event.On(save.EVNT_AccountLogout, AccountLogout)
}

func AccountLogout(args ...interface{}) {
	a := args[0].(*save.Account)
	log := logrus.WithFields(logrus.Fields{
		"Username": a.Username,
	})

	if err := accountlogout(a); err != nil {
		log.Error("cmd.AccountLogout: " + err.Error())
	} else {
		log.Warn("cmd.AccountLogout")
	}
}

func accountlogout(a *save.Account) error {
	if a == nil {
		return errors.New("Account already logout")
	}

	delete(save.Accounts, a.Username)

	if s := server.Sessions[a.SessionId]; s != nil {
		if err := sessionexpire(s); err != nil {
			return err
		}
	}
	if p := save.Players[a.Username]; p != nil {
		if err := playerdelete(p.Username); err != nil {
			return err
		} else if err := playerpush(p.Username); err != nil {
			return err
		}
	}

	return nil
}
