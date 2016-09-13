package cmd

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/engine"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/space"
	"github.com/ospokemon/ospokemon/util"
)

func init() {
	util.Event.On(save.EVNT_PlayerPull, PlayerPull)
}

func PlayerPull(args ...interface{}) {
	username := args[0].(string)

	p, err := playerpullall(username)

	if err != nil {
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

func playerpullall(username string) (*save.Player, error) {
	p := save.MakePlayer(username)

	if err := playerpull(p); err != nil {
		return nil, err
	}

	if err := playerpulllocation(p); err != nil {
		return nil, err
	}

	return p, nil
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
