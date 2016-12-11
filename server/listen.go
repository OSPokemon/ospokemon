package server

import (
	"github.com/ospokemon/ospokemon/save"
	"golang.org/x/net/websocket"
)

func Listen(s *Session) {
	for s.Websocket != nil {
		var message WebsocketMessage
		err := websocket.JSON.Receive(s.Websocket, &message)

		if err != nil {
			p, _ := save.GetPlayer(s.Username)
			location := p.Entity.Component(save.COMP_Location).(*save.Location)
			u, _ := save.GetUniverse(location.UniverseId)

			p.Entity.RemoveComponent(s)
			u.Remove(p.Entity)

			s.Websocket.Close()
			return
		}

		go ReceiveMessage(s, &message)
	}
}
