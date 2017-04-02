package persistence

import (
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"time"
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

	log.Add("Username", username).Info("accounts select")

	if player, err := ospokemon.GetPlayer(username); err == nil {
		player.AddPart(account)
		player.AddPart(player)
		account.Parts = player.Parts
	} else {
		return nil, err
	}

	return account, nil
}

func AccountsInsert(account *ospokemon.Account) error {
	_, err := Connection.Exec(
		"INSERT INTO accounts (username, password, register) values (?, ?, ?)",
		account.Username,
		account.Password,
		account.Register.Unix(),
	)

	if err == nil {
		log.Add("Username", account.Username).Info("accounts insert")

		if player := account.GetPlayer(); player != nil {
			ospokemon.Players.Insert(account.GetPlayer())
		}

		event.Fire(event.AccountsInsert, account)
	}

	return err
}

func AccountsDelete(account *ospokemon.Account) error {
	_, err := Connection.Exec("DELETE FROM accounts WHERE username=?", account.Username)

	if err == nil {
		log.Add("Username", account.Username).Info("accounts delete")
		ospokemon.Players.Delete(account.GetPlayer())
		event.Fire(event.AccountsDelete, account)
	}

	return err
}
