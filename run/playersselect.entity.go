package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
	"github.com/ospokemon/ospokemon/space"
)

func init() {
	event.On(event.PlayersSelect, PlayersSelectEntity)
}

func PlayersSelectEntity(args ...interface{}) {
	player := args[0].(*game.Player)
	entity, err := query.EntitiesPlayersSelect(player)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
			"Error":    err.Error(),
		}).Error("players select entity")
		return
	}

	if class, _ := query.GetClass(player.Class); class != nil {
		r := entity.Shape.(*space.Rect)
		r.Dimension.DX = class.Dimension.DX
		r.Dimension.DY = class.Dimension.DY
	}

	player.AddPart(entity)
	entity.Parts = player.Parts
}
