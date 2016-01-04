package loader

import (
	// "github.com/ospokemon/ospokemon/server"
	"errors"
)

func DoSignup(username string, password string, trainerclass int, speciesid int) error {
	existing_account := LoadAccount(username)

	if existing_account != nil {
		return errors.New("Signup fail. Account name already exists")
	}

	result, err := Connection.Exec("INSERT INTO players (name, password, class, x, y) VALUES (?, ?, ?, ?, ?);", username, password, trainerclass, 250, 250)

	if err != nil {
		return err
	}

	playerId, err := result.LastInsertId()

	if err != nil {
		return err
	}

	Connection.Exec("INSERT INTO players_stats (player_id, stat, value, maxvalue, basemaxvalue) VALUES (?, ?, ?, ?, ?);", playerId, "health", 100, 100, 100)
	Connection.Exec("INSERT INTO players_stats (player_id, stat, value, maxvalue, basemaxvalue) VALUES (?, ?, ?, ?, ?);", playerId, "speed", 25, 25, 25)

	return nil
}
