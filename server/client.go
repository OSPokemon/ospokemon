package server

import (
	"code.google.com/p/go.net/websocket"
	"github.com/cznic/mathutil"
)

var Clients = make(map[int]*Client)
var clientIdGen, _ = mathutil.NewFC32(0, 999999, true)

type Client struct {
	ClientId  int
	SessionId int
	Entities  []int
	Conn      *websocket.Conn
	Close     chan bool
}

var ConnectClient func(client *Client)
var DisconnectClient func(client *Client)
var ReceiveMessage func(c *Client, message map[string]interface{})

func CreateClient(conn *websocket.Conn, session *Session) *Client {
	if session == nil || session.Username == "" {
		return nil
	} else {
		session.ClientId = clientIdGen.Next()
	}

	Clients[session.ClientId] = &Client{session.ClientId, session.SessionId, make([]int, 0), conn, make(chan bool)}
	return Clients[session.ClientId]
}

func (c *Client) AddEntity(entityId interface{}) {
	c.Entities = append(c.Entities, entityId.(int))
}

func (c *Client) RemoveEntity(entityId interface{}) {
	for position := 0; position < len(c.Entities); position++ {
		if c.Entities[position] == entityId {
			c.Entities[position] = c.Entities[len(c.Entities)-1]
			c.Entities = c.Entities[:len(c.Entities)-1]
			return
		}
	}
}

func (c *Client) Send(message string) {
	websocket.Message.Send(c.Conn, message)
}
