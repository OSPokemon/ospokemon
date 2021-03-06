package persistence

import (
	"github.com/ospokemon/ospokemon"
	"github.com/ospokemon/ospokemon/space"
	"github.com/pkg/errors"
)

func init() {
	ospokemon.Players.Select = PlayersSelect
	ospokemon.Players.Insert = PlayersInsert
	ospokemon.Players.Delete = PlayersDelete
}

func PlayersSelect(username string) (*ospokemon.Player, error) {
	row := Connection.QueryRow(
		"SELECT level, experience, money, class, universe, x, y FROM players WHERE username=?",
		username,
	)

	var levelbuff, experiencebuff, moneybuff, classbuff uint
	entity := ospokemon.MakeEntity()
	r := entity.Shape.(*space.Rect)

	err := row.Scan(&levelbuff, &experiencebuff, &moneybuff, &classbuff, &entity.UniverseId, &r.Anchor.X, &r.Anchor.Y)
	if err != nil {
		return nil, errors.Wrap(err, "players.scan")
	}

	class, err := ospokemon.GetClass(classbuff)
	if err != nil {
		return nil, errors.Wrap(err, "players.getclass")
	}

	player := ospokemon.BuildPlayer(username, class, entity)
	player.Level = levelbuff
	player.Experience = experiencebuff
	player.Money = moneybuff

	actions, err := ActionsPlayersSelect(player)
	if err != nil {
		return nil, errors.Wrap(err, "players.partactions")
	}
	player.AddPart(actions)

	itembag, err := ItembagsPlayersSelect(player)
	if err != nil {
		return nil, errors.Wrap(err, "players.partitembag")
	}
	player.AddPart(itembag)

	stats, err := PlayersStatsSelect(player)
	if err != nil {
		return nil, errors.Wrap(err, "players.partstats")
	}
	player.AddPart(stats)

	err = BindingsPlayersSelect(player)
	if err != nil {
		return nil, errors.Wrap(err, "players.partbindings")
	}

	ospokemon.LOG().Add("Username", player.Username).Debug("players select")
	return player, nil
}

func PlayersInsert(player *ospokemon.Player) error {
	entity := player.GetEntity()
	r := entity.Shape.(*space.Rect)

	_, err := Connection.Exec(
		"INSERT INTO players (username, level, experience, money, class, universe, x, y) values (?, ?, ?, ?, ?, ?, ?, ?)",
		player.Username,
		player.Level,
		player.Experience,
		player.Money,
		player.Class,
		entity.UniverseId,
		r.Anchor.X,
		r.Anchor.Y,
	)
	if err != nil {
		return errors.New("players insert: " + err.Error())
	}

	err = BindingsPlayersInsert(player)
	if err != nil {
		return err
	}

	err = ActionsPlayersInsert(player)
	if err != nil {
		return err
	}

	err = ItembagsPlayersInsert(player)
	if err != nil {
		return err
	}

	err = PlayersStatsInsert(player)
	if err != nil {
		return err
	}

	ospokemon.LOG().Add("Username", player.Username).Debug("players insert")
	return err
}

func PlayersDelete(player *ospokemon.Player) error {
	_, err := Connection.Exec("DELETE FROM players WHERE username=?", player.Username)
	if err != nil {
		return err
	}

	err = ActionsPlayersDelete(player)
	if err != nil {
		return err
	}

	err = PlayersStatsDelete(player)
	if err != nil {
		return err
	}

	err = BindingsPlayersDelete(player)
	if err != nil {
		return err
	}

	err = ItembagsPlayersDelete(player)
	if err != nil {
		return err
	}

	ospokemon.LOG().Add("Username", player.Username).Debug("players delete")
	return nil
}
