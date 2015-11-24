package connection

import (
	"encoding/json"
	"github.com/ospokemon/ospokemon/world"
	"log"
	"strconv"
	"time"
)

func Update(base map[string]*world.BasicView, now time.Time) {
	for _, client := range Clients {
		view := make(map[string]interface{})
		view["world"] = copyMap(base)

		controlViews := make(map[string]*world.FullView)
		for _, id := range client.Entities {
			tag := strconv.Itoa(id)
			controlViews[tag] = world.MakeFullView(id, world.Entities[id], now)
		}
		view["control"] = controlViews

		json, _ := json.Marshal(view)
		message := string(json)

		client.Send <- message
	}
}

func ReceiveMessage(name string, message map[string]interface{}) {
	client := Clients[name]

	entityId := int(message["entity"].(float64))

	var entity world.Entity
	for _, id := range client.Entities {
		if id == entityId {
			entity = world.Entities[id]
		}
	}

	if entity == nil {
		return
	}

	log.Printf("message received: %v", message)

	if message["walk"] != nil {
		coords := message["walk"].(map[string]interface{})
		walking := &world.Position{}
		walking.X = coords["x"].(float64)
		walking.Y = coords["y"].(float64)
		entity.Physics().Walking = walking
	} else if message["ability"] != nil {
		ability := message["ability"].(string)
		action := &world.Action{}

		switch target := message["target"].(type) {
		default:
			break
		case map[string]interface{}:
			action.Target = &world.Position{
				X: target["x"].(float64),
				Y: target["y"].(float64),
			}
			break
		case int:
			action.Target = target
			break
		}

		log.Printf("Action accepted for client(%s)(%d): %v", client.Name, entityId, action)

		action.Ability = entity.Controls().Abilities[ability]
		entity.Controls().Action = action
	}
}

func copyMap(src map[string]*world.BasicView) map[string]*world.BasicView {
	dst := make(map[string]*world.BasicView)

	for k, v := range src {
		dst[k] = v
	}

	return dst
}
