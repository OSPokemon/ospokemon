package run

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
)

func init() {
	event.On(event.PlayersSelect, PlayersSelectImaging)
}

func PlayersSelectImaging(args ...interface{}) {
	player := args[0].(*game.Player)
	class, err := query.GetClass(player.Class)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Username": player.Username,
			"Class":    player.Class,
			"Error":    err.Error(),
		}).Error("players select imaging")
		return
	}

	imaging := game.MakeImaging()
	imaging.ReadAnimations(class.Animations)

	player.AddPart(imaging)
}
