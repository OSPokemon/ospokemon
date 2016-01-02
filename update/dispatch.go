package update

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/physics"
	"github.com/ospokemon/ospokemon/server"
	"github.com/ospokemon/ospokemon/world"
)

func ReceiveMessage(client *server.Client, message map[string]interface{}) {
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
		"Client":  client.ClientId,
		"message": message,
	}).Debug("WS received")

	if message["walk"] != nil {
		point := parsePoint(message["walk"])
		entity.SetWalking(&point)
	} else if message["ability"] != nil {
		ability := message["ability"].(string)
		action := &world.Action{
			Ability: entity.Abilities()[ability],
		}

		switch target := message["target"].(type) {
		default:
			log.WithFields(log.Fields{
				"client": client.ClientId,
				"target": target,
			}).Warn("Message received with unrecognized target type")
			break
		case map[string]interface{}:
			action.Target = parsePoint(target)
			break
		case float64:
			action.Target = world.Entities[int(target)]
			break
		}

		entity.SetAction(action)

		log.WithFields(log.Fields{
			"client": client.ClientId,
			"entity": entityId,
			"action": action,
		}).Debug("Action accepted for client")
	} else {
		log.WithFields(log.Fields{
			"Client":   client.ClientId,
			"EntityId": entityId,
			"Message":  message,
		}).Warn("Unrecognized message format")
	}
}

func parsePoint(src interface{}) physics.Point {
	coords := src.(map[string]interface{})
	point := physics.Point{}
	point.X = coords["x"].(float64)
	point.Y = coords["y"].(float64)
	return point
}
