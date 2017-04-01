package server

import (
	"golang.org/x/net/websocket"
	"ospokemon.com"
	"ospokemon.com/log"
	"ospokemon.com/option"
	"ospokemon.com/query"
)

func Listen(s *Session) {
	for s.Websocket != nil {
		var message WebsocketMessage
		err := websocket.JSON.Receive(s.Websocket, &message)

		if err != nil {
			if err.Error() != "EOF" {
				log.Warn(err.Error())
			}

			if account := ospokemon.Accounts[s.Username]; !option.Bool("allow-refresh") && account != nil {
				query.AccountsDelete(account)
				query.AccountsInsert(account)
			}

			s.Websocket.Close()

			if player := ospokemon.Players[s.Username]; player != nil {
				entity := player.GetEntity()
				universe := ospokemon.Multiverse[entity.UniverseId]

				universe.Remove(entity)
			}

			return
		}

		go ReceiveMessage(s, &message)
	}
}
