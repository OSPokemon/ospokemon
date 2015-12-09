package connection

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/world"
)

func ReceiveMessage(name string, message map[string]interface{}) {
	client := Clients[name]

	entityId := int(message["entity"].(float64))

	var entity world.Intelligence
	for _, id := range client.Entities {
		if id == entityId {
			entity = world.Entities[id].(world.Intelligence)
		}
	}

	if entity == nil {
		return
	}

	log.WithFields(log.Fields{
		"Client":  name,
		"message": message,
	}).Debug("WS received")

	if message["walk"] != nil {
		coords := message["walk"].(map[string]interface{})
		walking := &world.Position{}
		walking.X = coords["x"].(float64)
		walking.Y = coords["y"].(float64)
		entity.SetWalking(walking)
	} else if message["ability"] != nil {
		ability := message["ability"].(string)
		action := &world.Action{
			Ability: entity.Abilities()[ability],
		}

		switch target := message["target"].(type) {
		default:
			log.WithFields(log.Fields{
				"client": client.Name,
				"target": target,
			}).Warn("Message received with unrecognized target type")
			break
		case map[string]interface{}:
			action.Target = &world.Position{
				X: target["x"].(float64),
				Y: target["y"].(float64),
			}
			break
		case float64:
			action.Target = world.Entities[int(target)]
			break
		}

		entity.SetAction(action)

		log.WithFields(log.Fields{
			"client": client.Name,
			"entity": entityId,
			"action": action,
		}).Debug("Action accepted for client")
	} else {
		log.WithFields(log.Fields{
			"Client":   client.Name,
			"EntityId": entityId,
			"Message":  message,
		}).Warn("Unrecognized message format")
	}
}

// target data can be coppied

type viewcopier map[string]interface{}

func (src viewcopier) copy() map[string]interface{} {
	dst := make(map[string]interface{})

	for k, v := range src {
		dst[k] = v
	}

	return dst
}
