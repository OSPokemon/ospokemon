package connection

import (
	"code.google.com/p/go.net/websocket"
)

var Clients = make(map[string]*Client)

type Client struct {
	Name     string
	Entities []int
	Conn     *websocket.Conn
	Send     chan string
	Close    chan bool
}

func NewClient(name string, conn *websocket.Conn) *Client {
	return &Client{
		Name:     name,
		Conn:     conn,
		Entities: make([]int, 0),
		Send:     make(chan string, 1),
		Close:    make(chan bool, 1),
	}
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
