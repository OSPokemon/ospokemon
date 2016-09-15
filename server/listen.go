package server

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/util"
	"golang.org/x/net/websocket"
)

func Listen(s *Session) {
	for s.Websocket != nil {
		var message WebsocketMessage
		err := websocket.JSON.Receive(s.Websocket, &message)

		if err != nil {
			if s.Websocket != nil {
				logrus.Warn("server.Listen: " + err.Error())
				util.Event.Fire(EVNT_WebsocketDisconnect, s)
			}

			return
		} else {
			go util.Event.Fire(EVNT_WebsocketMessage, s, message)
		}
	}
}
