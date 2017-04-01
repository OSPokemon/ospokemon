package query

import (
	"ospokemon.com/event"
	"ospokemon.com/game"
	"ospokemon.com/log"
	"ospokemon.com/space"
)

func PlayersInsert(player *game.Player) error {
	entity := player.GetEntity()
	r := entity.Shape.(*space.Rect)
	_, err := Connection.Exec(
		"INSERT INTO players (username, level, experience, money, class, bagsize, universe, x, y) values (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		player.Username,
		player.Level,
		player.Experience,
		player.Money,
		player.Class,
		player.BagSize,
		entity.UniverseId,
		r.Anchor.X,
		r.Anchor.Y,
	)

	if err == nil {
		log.Add("Username", player.Username).Info("players insert")

		delete(game.Players, player.Username)
		event.Fire(event.PlayersInsert, player)
	}

	return err
}
