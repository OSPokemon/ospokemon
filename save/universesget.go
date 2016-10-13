package save

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/engine"
)

func UniversesGet(id uint) (*engine.Universe, error) {
	u := engine.MakeUniverse(id)
	row := Connection.QueryRow(
		"SELECT x, y, dx, dy FROM universes WHERE id=?",
		id,
	)

	if err := row.Scan(&u.Space.Rect.Anchor.X, &u.Space.Rect.Anchor.Y, &u.Space.Rect.Dimension.DX, &u.Space.Rect.Dimension.DY); err != nil {
		return nil, errors.New("universeget: " + err.Error())
	}

	logrus.WithFields(logrus.Fields{
		"Id": id,
	}).Debug("save.UniversesGet")

	return u, nil
}
