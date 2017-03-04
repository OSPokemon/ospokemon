package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/part"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.PlayersInsert, PlayersInsertEntity)
}

func PlayersInsertEntity(args ...interface{}) {
	player := args[0].(*game.Player)
	entity, ok := player.Parts[part.Entity].(*game.Entity)

	if !ok {
		entity = game.MakeEntity()
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
		}).Debug("players insert entity: grant empty entity")
	}

	err := query.EntitiesPlayersInsert(player, entity)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
			"Error":    err.Error(),
		}).Error("players insert entity")
	}
}
