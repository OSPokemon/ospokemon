package cmd

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/space"
	"github.com/ospokemon/ospokemon/util"
	"time"
)

func init() {
	util.Event.On(save.EVNT_PlayerPull, PlayerPull)
}

func PlayerPull(args ...interface{}) {
	username := args[0].(string)
	p := save.MakePlayer(username)

	if err := playerpullall(p); err != nil {
		logrus.WithFields(logrus.Fields{
			"Username": username,
		}).Error("cmd.PlayerPull: " + err.Error())
		return
	}

	save.Players[username] = p

	logrus.WithFields(logrus.Fields{
		"Username": username,
	}).Info("cmd.PlayerPull")
}

func playerpullall(p *save.Player) error {
	if err := playerpull(p); err != nil {
		return err
	}
	if err := playerpulllocation(p); err != nil {
		return err
	}
	if err := playerpullbindings(p); err != nil {
		return err
	}

	return nil
}

func playerpull(p *save.Player) error {
	row := save.Connection.QueryRow(
		"SELECT level, experience, money FROM players WHERE username=?",
		p.Username,
	)

	if err := row.Scan(&p.Level, &p.Experience, &p.Money); err != nil {
		return errors.New("queryplayer: " + err.Error())
	}

	return nil
}

func playerpulllocation(p *save.Player) error {
	row := save.Connection.QueryRow(
		"SELECT universe, x, y, dx, dy FROM locations_players WHERE username=?",
		p.Username,
	)

	l := p.Entity.Component(engine.COMP_Location).(*engine.Location)
	r := l.Shape.(*space.Rect)

	if err := row.Scan(&l.UniverseId, &r.Anchor.X, &r.Anchor.Y, &r.Dimension.DX, &r.Dimension.DY); err != nil {
		return errors.New("queryplayerlocation: " + err.Error())
	}

	return nil
}

func playerpullactions(p *save.Player) error {
	rows, err := save.Connection.Query(
		"SELECT spellid, timer FROM actions_players WHERE username=?",
		p.Username,
	)

	if err != nil {
		return err
	} else {
		defer rows.Close()
	}

	a := p.Entity.Component(engine.COMP_Actions).(engine.Actions)

	for rows.Next() {
		var spellidbuff uint
		var timebuff uint64
		err = rows.Scan(spellidbuff, timebuff)

		if err != nil {
			return err
		}

		action := engine.MakeAction(spellidbuff)

		if timebuff > 0 {
			t := time.Duration(timebuff)
			action.Timer = &t
		}

		a[action.SpellId] = action
	}

	return nil
}

func playerpullbindings(p *save.Player) error {
	rows, err := save.Connection.Query(
		"SELECT key, spellid FROM bindings_players WHERE username=?",
		p.Username,
	)

	if err != nil {
		return err
	} else {
		defer rows.Close()
	}

	a := p.Entity.Component(engine.COMP_Actions).(engine.Actions)
	b := p.Entity.Component(engine.COMP_Bindings).(engine.Bindings)

	for rows.Next() {
		var keybuff string
		var spellidbuff uint
		binding := &engine.Binding{}
		err = rows.Scan(&keybuff, &spellidbuff)

		if err != nil {
			return err
		}

		binding.Action = a[spellidbuff]
		b[keybuff] = binding
	}

	return nil
}
