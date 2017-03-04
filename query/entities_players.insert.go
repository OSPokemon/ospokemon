package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/space"
)

func EntitiesPlayersInsert(player *game.Player, entity *game.Entity) error {
	r := entity.Shape.(*space.Rect)

	_, err := Connection.Exec(
		"INSERT INTO entities_players (username, universe, x, y) VALUES (?, ?, ?, ?)",
		player.Username,
		entity.UniverseId,
		r.Anchor.X,
		r.Anchor.Y,
	)

	if err == nil {
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
			"Universe": entity.UniverseId,
			"Locatoin": r.Anchor,
		}).Debug("entities_players insert")
	}

	return err
}
