package persistence

import (
	"errors"
	"time"

	"ospokemon.com"
	"ztaylor.me/log"
)

func init() {
	ospokemon.Accounts.Select = AccountsSelect
	ospokemon.Accounts.Insert = AccountsInsert
	ospokemon.Accounts.Delete = AccountsDelete
}

func AccountsSelect(username string) (*ospokemon.Account, error) {
	row := Connection.QueryRow(
		"SELECT password, register FROM accounts WHERE username=?",
		username,
	)

	account := ospokemon.MakeAccount(username)
	var timebuff int64
	err := row.Scan(&account.Password, &timebuff)

	if err != nil {
		return nil, err
	}

	account.Register = time.Unix(timebuff, 0)

	if player, err := ospokemon.GetPlayer(username); err == nil {
		player.AddPart(account)
		player.AddPart(player)
		account.Parts = player.Parts
	} else {
		return nil, err
	}

	log.Add("Username", username).Debug("accounts select")
	return account, nil
}

func AccountsInsert(account *ospokemon.Account) error {
	_, err := Connection.Exec(
		"INSERT INTO accounts (username, password, register) values (?, ?, ?)",
		account.Username,
		account.Password,
		account.Register.Unix(),
	)

	if err != nil {
		return errors.New("accounts insert: " + err.Error())
	}

	err = ospokemon.Players.Insert(account.GetPlayer())
	if err != nil {
		return err
	}

	log.Add("Username", account.Username).Debug("accounts insert")
	return nil
}

func AccountsDelete(account *ospokemon.Account) error {
	_, err := Connection.Exec("DELETE FROM accounts WHERE username=?", account.Username)
	if err != nil {
		return err
	}

	err = ospokemon.Players.Delete(account.GetPlayer())
	if err != nil {
		return errors.New("accounts delete: " + err.Error())
	}

	log.Add("Username", account.Username).Debug("accounts delete")
	return nil
}
