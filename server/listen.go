package server

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/query"
	"golang.org/x/net/websocket"
)

func Listen(s *Session) {
	for s.Websocket != nil {
		var message WebsocketMessage
		err := websocket.JSON.Receive(s.Websocket, &message)

		if err != nil {
			if err.Error() != "EOF" {
				logrus.Warn(err.Error())
			}

			account := game.Accounts[s.Username]

			if account != nil {
				query.AccountsDelete(account)
				query.AccountsInsert(account)
			}

			s.Websocket.Close()

			if player := game.Players[s.Username]; player != nil {
				entity := player.GetEntity()
				universe := game.Multiverse[entity.UniverseId]

				universe.Remove(entity)
			}

			return
		}

		go ReceiveMessage(s, &message)
	}
}
