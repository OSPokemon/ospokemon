package server

import (
	"code.google.com/p/go.net/websocket"
	log "github.com/Sirupsen/logrus"
	"strconv"
)

var WebsocketHandler = websocket.Handler(func(conn *websocket.Conn) {
	var session *Session

	if sessionCookie, err := conn.Request().Cookie("SessionId"); err == nil {
		if sessionId, err := strconv.ParseInt(sessionCookie.Value, 10, 0); err == nil {
			session = Sessions[int(sessionId)]
		}
	}

	if session == nil {
		log.Info("Websocket connection rejected, Session missing")
		return
	}

	account := Accounts[session.Username]
	if account == nil {
		log.WithFields(log.Fields{
			"Session": session,
		}).Info("Invalid session detected; destroy session")
		delete(Sessions, session.SessionId)
		return
	}

	client := CreateClient(conn, session)

	log.WithFields(log.Fields{
		"Session": client.SessionId,
		"Client":  client.ClientId,
	}).Info("Client created")

	ConnectClient(client)
	listenClient(client)
	DisconnectClient(client)

	log.WithFields(log.Fields{
		"Session": client.SessionId,
		"Client":  client.ClientId,
	}).Info("Client closed")

	delete(Clients, session.ClientId)
	conn.Close()
})

func listenClient(c *Client) {
	for {
		var message map[string]interface{}
		err := websocket.JSON.Receive(c.Conn, &message)

		if err != nil {
			return
		} else {
			go ReceiveMessage(c, message)
		}
	}
}
