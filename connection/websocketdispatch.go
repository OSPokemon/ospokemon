package connection

import (
	"encoding/json"
	"github.com/ospokemon/ospokemon/world"
	"log"
	"strconv"
	"time"
)

func Update(base map[string]*world.View) {
	for _, client := range Clients {
		view := make(map[string]interface{})
		view["world"] = copyMap(base)

		controlViews := make(map[string]*world.View)
		for _, id := range client.Entities {
			tag := strconv.Itoa(id)
			controlViews[tag] = base[tag]
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
	ability := message["ability"].(string)

	var entity world.Entity
	for _, id := range client.Entities {
		if id == entityId {
			entity = world.Entities[id]
		}
	}

	if entity == nil {
		return
	}
	if entity.Controls().State&world.CTRLPnocast > 0 {
		return
	}
	if entity.Controls().State&world.CTRLPstuck > 0 && ability == "walk" {
		return
	}

	action := &world.Action{}
	action.Clock = time.Now()

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

	if ability == "walk" {
		action.Ability = world.WalkAbility
	} else {
		action.Ability = entity.Controls().Abilities[ability]
	}

	log.Printf("Action accepted for client(%s)(%d): %v", client.Name, entityId, action)

	entity.Controls().Action = action
}

func copyMap(src map[string]*world.View) map[string]*world.View {
	dst := make(map[string]*world.View)

	for k, v := range src {
		dst[k] = v
	}

	return dst
}
