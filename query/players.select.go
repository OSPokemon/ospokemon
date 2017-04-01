package query

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"ospokemon.com/space"
)

func PlayersSelect(username string) (*ospokemon.Player, error) {
	row := Connection.QueryRow(
		"SELECT level, experience, money, class, bagsize, universe, x, y FROM players WHERE username=?",
		username,
	)

	var levelbuff, experiencebuff, moneybuff, classbuff, bagsizebuff uint

	entity := ospokemon.MakeEntity()
	r := entity.Shape.(*space.Rect)

	err := row.Scan(&levelbuff, &experiencebuff, &moneybuff, &classbuff, &bagsizebuff, &entity.UniverseId, &r.Anchor.X, &r.Anchor.Y)
	if err != nil {
		return nil, err
	}

	class, err := GetClass(classbuff)
	if err != nil {
		return nil, err
	}

	player := ospokemon.BuildPlayer(username, bagsizebuff, class, entity)
	player.Level = levelbuff
	player.Experience = experiencebuff
	player.Money = moneybuff

	ospokemon.Players[username] = player

	log.Add("Username", player.Username).Info("players select")

	event.Fire(event.PlayersSelect, player)

	return player, nil
}
