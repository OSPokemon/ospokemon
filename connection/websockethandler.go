package connection

import (
	"code.google.com/p/go.net/websocket"
	"github.com/ospokemon/ospokemon/data"
	"github.com/ospokemon/ospokemon/world"
	"log"
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
	client.Entities = data.FullLoadPlayer(name)
	Clients[name] = client

	log.Printf("WSclient connection created as %s: %v", name, client.Entities)

	go client.ListenSend()
	client.ListenRead()

	data.FullUnloadPlayer(name)
	for _, id := range client.Entities {
		world.RemoveEntity(id)
	}

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

// func PushWorld(t time.Time) {
// 	message := make(map[string]interface{})

// 	world.Entities.Lock()
// 	for entityId, dirty := range world.Entities.Dirty {
// 		if !dirty {
// 			continue
// 		}

// 		entity := world.Entities.All[entityId]
// 		entityJson := make(map[string]interface{})

// 		entityJson["tag"] = entity.Tag()
// 		entityJson["x"] = entity.Position().X
// 		entityJson["y"] = entity.Position().Y

// 		if entity.Action() != nil {
// 			action := entity.Action()

// 			if action.Type == world.CastAction {
// 				castLength := t.Sub(action.Start)
// 				castPercentage := int(float64(castLength) / float64(action.Duration) * 100)
// 				entityJson["cast"] = castPercentage
// 			} else if action.Type == world.MoveAction {
// 				entityJson["move"] = action.Data.(world.Position)
// 			}
// 		}

// 		message[entity.Tag()] = entityJson
// 	}
// 	world.Entities.Dirty = make(map[string]bool)
// 	world.Entities.Unlock()

// 	for _, client := range clients {
// 		client.Send <- &message
// 	}
// }
