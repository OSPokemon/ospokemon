package db

import (
	"github.com/ospokemon/ospokemon/server"
	"time"
)

func LoadAccount(username string) (*server.Account, error) {
	row := Connection.QueryRow("SELECT username, password, register FROM accounts WHERE username=?", username)

	account := &server.Account{
		TrainerIds:  make([]int, 0),
		Permissions: make(map[string]bool),
	}
	t := 0

	err := row.Scan(&account.Username, &account.Password, &t)
	if err != nil {
		return nil, err
	}
	account.Register = time.Unix(int64(t), 0)

	rows, err := Connection.Query("SELECT id FROM trainers WHERE username=?", username)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var TrainerId int
		rows.Scan(&TrainerId)
		account.TrainerIds = append(account.TrainerIds, TrainerId)
	}

	rows, err = Connection.Query("SELECT permission FROM accounts_permissions WHERE username=?", username)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var permission string
		rows.Scan(&permission)
		account.Permissions[permission] = true
	}

	return account, nil
}
