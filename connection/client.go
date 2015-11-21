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
		Entities: make([]int, 1),
		Send:     make(chan string, 1),
		Close:    make(chan bool, 1),
	}
}
