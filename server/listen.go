package server

import (
	"golang.org/x/net/websocket"
	"ospokemon.com"
	"ospokemon.com/log"
	"ospokemon.com/option"
)

func Listen(s *Session) {
	for s.Websocket != nil {
		var message WebsocketMessage
		err := websocket.JSON.Receive(s.Websocket, &message)

		if err != nil {
			if err.Error() != "EOF" {
				log.Warn(err.Error())
			}

			account, _ := ospokemon.GetAccount(s.Username)
			if !option.Bool("allow-refresh") && account != nil && account.Parts[PARTsession] != nil {
				ospokemon.Accounts.Delete(account)
				ospokemon.Accounts.Insert(account)
			}

			s.Websocket.Close()

			if player, _ := ospokemon.GetPlayer(s.Username); player != nil {
				entity := player.GetEntity()
				universe := ospokemon.Multiverse[entity.UniverseId]

				universe.Remove(entity)
			}

			return
		}

		go ReceiveMessage(s, &message)
	}
}
