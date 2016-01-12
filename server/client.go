package server

import (
	"code.google.com/p/go.net/websocket"
	"github.com/cznic/mathutil"
)

type Client struct {
	ClientId  int
	SessionId int
	Entities  []int
	Conn      *websocket.Conn
}

var Clients = make(map[int]*Client)
var clientIdGen, _ = mathutil.NewFC32(1, 999999999, true)

var ConnectClient func(client *Client)
var DisconnectClient func(client *Client)
var ReceiveMessage func(username string, message map[string]interface{})

func CreateClient(conn *websocket.Conn) *Client {
	return &Client{
		ClientId: clientIdGen.Next(),
		Entities: make([]int, 0),
		Conn:     conn,
	}
}

func (c *Client) AddEntity(entityId interface{}) {
	c.Entities = append(c.Entities, entityId.(int))
}

func (c *Client) RemoveEntity(entityId interface{}) {
	for i := 0; i < len(c.Entities); i++ {
		if c.Entities[i] == entityId {
			c.Entities[i] = c.Entities[len(c.Entities)-1]
			c.Entities = c.Entities[:len(c.Entities)-1]
			return
		}
	}
}

func (c *Client) Send(message string) {
	websocket.Message.Send(c.Conn, message)
}
