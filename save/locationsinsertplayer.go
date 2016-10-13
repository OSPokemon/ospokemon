package save

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/space"
)

func LocationsInsertPlayer(username string, l *Location) error {
	r := l.Shape.(*space.Rect)

	_, err := Connection.Exec(
		"INSERT INTO locations_players (username, universe, x, y, dx, dy) VALUES (?, ?, ?, ?, ?, ?)",
		username,
		l.UniverseId,
		r.Anchor.X,
		r.Anchor.Y,
		r.Dimension.DX,
		r.Dimension.DY,
	)

	if err != nil {
		return errors.New("locationsinsertplayer: " + err.Error())
	}

	logrus.WithFields(logrus.Fields{
		"Username": username,
		"Location": l,
	}).Debug("save.LocationsInsertPlayer")

	return nil
}
