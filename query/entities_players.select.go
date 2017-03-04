package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/space"
)

func EntitiesPlayersSelect(player *game.Player) (*game.Entity, error) {
	row := Connection.QueryRow(
		"SELECT universe, x, y FROM entities_players WHERE username=?",
		player.Username,
	)

	entity := game.MakeEntity()
	r := entity.Shape.(*space.Rect)
	err := row.Scan(&entity.UniverseId, &r.Anchor.X, &r.Anchor.Y)

	if err == nil {
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
			"Entity":   entity,
		}).Debug("entities_players select")
	}

	return entity, err
}
