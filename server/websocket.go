package server

import (
	"code.google.com/p/go.net/websocket"
	log "github.com/Sirupsen/logrus"
)

var WebsocketHandler = websocket.Handler(func(conn *websocket.Conn) {
	sessionId := readSessionId(conn.Request())
	session := Sessions[sessionId]

	if session == nil {
		log.Info("Websocket connection rejected; Session missing")
		return
	}

	if session.ClientId > 0 {
		log.Info("Websocket connection rejected; Websocket already open for session")
		return
	}

	account := Accounts[session.Username]

	if account == nil {
		log.Info("Websocket connection rejected: Account missing")
		return
	}

	if account.TrainerId < 1 {
		log.Info("Websocket connection rejected: TrainerId missing")
		return
	}

	client := CreateClient(conn)

	session.ClientId = client.ClientId
	client.SessionId = session.SessionId
	Clients[client.ClientId] = client

	ConnectClient(client)
	listenDispatch(client.Conn, session.Username)
	DisconnectClient(client)

	log.WithFields(log.Fields{
		"Session": client.SessionId,
		"Client":  client.ClientId,
	}).Info("Client closed")

	delete(Clients, session.ClientId)
	session.ClientId = 0
	conn.Close()
})

func listenDispatch(conn *websocket.Conn, username string) {
	for {
		var message map[string]interface{}
		err := websocket.JSON.Receive(conn, &message)

		if err != nil {
			return
		} else {
			go ReceiveMessage(username, message)
		}
	}
}
