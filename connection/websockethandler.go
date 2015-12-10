package connection

import (
	"code.google.com/p/go.net/websocket"
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/objects/entities"
	"github.com/ospokemon/ospokemon/registry"
	"github.com/ospokemon/ospokemon/world"
	"strconv"
)

var ConnectHandler = websocket.Handler(func(conn *websocket.Conn) {
	var name string

	if sessionCookie, err := conn.Request().Cookie("SessionId"); err == nil {
		sessionId, err := strconv.ParseInt(sessionCookie.Value, 10, 0)

		if err == nil {
			name = Sessions[int(sessionId)]
		}
	}

	if name == "" {
		return
	}

	client := NewClient(name, conn)
	Clients[name] = client

	account := registry.Accounts[name]

	registry.Loaders["Player"](account.PlayerId)
	player := registry.Players[account.PlayerId]
	world.AddEntity(player)

	go preloadPokemon(player)

	client.AddEntity(player.EntityId())
	for _, pokemonId := range player.Pokemon() {
		if registry.Pokemon[pokemonId] != nil && registry.Pokemon[pokemonId].EntityId() > 0 {
			client.AddEntity(registry.Pokemon[pokemonId].EntityId())
		}
	}

	player.On("SummonPokemon", client.AddEntity)
	player.On("DismissPokemon", client.RemoveEntity)

	account.Online = true

	log.WithFields(log.Fields{
		"name":      name,
		"EntityIDs": client.Entities,
	}).Info("WSclient connection created")

	go client.ListenSend()
	client.ListenRead()

	delete(Clients, name)
	conn.Close()
})

func (c *Client) ListenSend() {
	for {
		select {
		case <-c.Close:
			c.Close <- true // echo
			return
		case message := <-c.Send:
			go websocket.Message.Send(c.Conn, message)
			break
		}
	}
}

func (c *Client) ListenRead() {
	for {
		select {
		case <-c.Close:
			c.Close <- true // echo
			return
		default:
			var message map[string]interface{}
			err := websocket.JSON.Receive(c.Conn, &message)

			if err != nil {
				log.Printf("WSclient connection closed(%v):%s", err, c.Name)
				c.Close <- true
			} else {
				go ReceiveMessage(c.Name, message)
			}
		}
	}
}

func preloadPokemon(player *entities.Player) {
	for _, pokemonId := range player.Pokemon() {
		registry.Loaders["Pokemon"](pokemonId)
	}
}
