package save

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/space"
)

func LocationsGetPlayer(username string) (*Location, error) {
	r := &space.Rect{
		Anchor:    space.Point{},
		Dimension: space.Vector{},
	}
	l := &Location{
		Shape: r,
	}

	row := Connection.QueryRow(
		"SELECT universe, x, y, dx, dy FROM locations_players WHERE username=?",
		username,
	)

	if err := row.Scan(&l.UniverseId, &r.Anchor.X, &r.Anchor.Y, &r.Dimension.DX, &r.Dimension.DY); err != nil {
		return nil, errors.New("locationsgetplayer: " + err.Error())
	}

	logrus.WithFields(logrus.Fields{
		"Username": username,
		"Location": l,
	}).Debug("save.LocationsGetPlayer")

	return l, nil
}
