package server

import (
	"github.com/ospokemon/ospokemon/part"
	"github.com/ospokemon/ospokemon/save"
	"golang.org/x/net/websocket"
)

func Listen(s *Session) {
	for s.Websocket != nil {
		var message WebsocketMessage
		err := websocket.JSON.Receive(s.Websocket, &message)

		if err != nil {
			p, _ := save.GetPlayer(s.Username)
			e := p.Parts[part.ENTITY].(*save.Entity)
			u, _ := save.GetUniverse(e.UniverseId)

			e.RemovePart(s)
			u.Remove(e)

			s.Websocket.Close()
			return
		}

		go ReceiveMessage(s, &message)
	}
}
