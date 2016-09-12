package server

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/util"
	"golang.org/x/net/websocket"
)

const EVNT_WebsocketConnect = "server/Websocket.Connect"
const EVNT_WebsocketMessage = "server/Websocket.Message"
const EVNT_WebsocketDisconnect = "server/Websocket.Disconnect"

var WebsocketHandler = websocket.Handler(func(conn *websocket.Conn) {
	s := readsession(conn.Request())

	if s == nil {
		logrus.Warn("server.WebsocketHandler: Failure: Session missing")
		return
	}

	if s.Websocket != nil {
		logrus.Warn("server.WebsocketHandler: Failure: Websocket already connected")
		return
	}

	s.Websocket = conn

	util.Event.Fire(EVNT_WebsocketConnect, s)

	listenDispatch(s)
})

func listenDispatch(s *Session) {
	for {
		var message WebsocketMessage
		err := websocket.JSON.Receive(s.Websocket, &message)

		if err != nil {
			util.Event.Fire(EVNT_WebsocketDisconnect, s, err.Error())
			return
		} else {
			go util.Event.Fire(EVNT_WebsocketMessage, s, message)
		}
	}
}
