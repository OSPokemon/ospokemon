package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/space"
	"github.com/ospokemon/ospokemon/util"
)

func init() {
	util.Event.On(save.EVNT_PlayerPush, PlayerPush)
}

func PlayerPush(args ...interface{}) {
	username := args[0].(string)
	p := save.Players[username]

	if err := playerpushall(p); err != nil {
		logrus.Error("cmd.PlayerPush: " + err.Error())
		return
	}

	logrus.WithFields(map[string]interface{}{
		"Username": p.Username,
	}).Warn("cmd.PlayerPush")
}

func playerpushall(p *save.Player) error {
	if err := playerpush(p); err != nil {
		return err
	}
	if err := playerpushlocation(p); err != nil {
		return err
	}
	if err := playerpushbindings(p); err != nil {
		return err
	}

	return nil
}

func playerpush(p *save.Player) error {
	_, err := save.Connection.Exec(
		"INSERT INTO players (username, level, experience, money) values (?, ?, ?, ?)",
		p.Username,
		p.Level,
		p.Experience,
		p.Money,
	)

	return err
}

func playerpushlocation(p *save.Player) error {
	l := p.Entity.Component(engine.COMP_Location).(*engine.Location)
	r := l.Shape.(*space.Rect)

	_, err := save.Connection.Exec(
		"INSERT INTO locations_players (username, universe, x, y, dx, dy) VALUES (?, ?, ?, ?, ?, ?)",
		p.Username,
		l.UniverseId,
		r.Anchor.X,
		r.Anchor.Y,
		r.Dimension.DX,
		r.Dimension.DY,
	)

	return err
}

func playerpushbindings(p *save.Player) error {
	b := p.Entity.Component(engine.COMP_Bindings).(engine.Bindings)

	for key, action := range b {
		_, err := save.Connection.Exec(
			"INSERT INTO bindings_players (username, key, name, image, script, casttime, cooldown) VALUES (?, ?, ?, ?, ?, ?, ?)",
			p.Username,
			key,
			action.Name,
			action.Image,
			action.ScriptId,
			action.CastTime,
			action.Cooldown,
		)

		if err != nil {
			return err
		}
	}

	return nil
}
