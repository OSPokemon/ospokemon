package query

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/space"
)

func PlayersSelect(username string) (*game.Player, error) {
	row := Connection.QueryRow(
		"SELECT level, experience, money, class, bagsize, universe, x, y FROM players WHERE username=?",
		username,
	)

	var levelbuff, experiencebuff, moneybuff, classbuff, bagsizebuff uint

	entity := game.MakeEntity()
	r := entity.Shape.(*space.Rect)

	err := row.Scan(&levelbuff, &experiencebuff, &moneybuff, &classbuff, &bagsizebuff, &entity.UniverseId, &r.Anchor.X, &r.Anchor.Y)
	if err != nil {
		return nil, err
	}

	class, err := GetClass(classbuff)
	if err != nil {
		return nil, err
	}

	player := game.BuildPlayer(username, bagsizebuff, class, entity)
	player.Level = levelbuff
	player.Experience = experiencebuff
	player.Money = moneybuff

	game.Players[username] = player

	logrus.WithFields(logrus.Fields{
		"Username": player.Username,
	}).Info("players select")

	event.Fire(event.PlayersSelect, player)

	return player, nil
}
