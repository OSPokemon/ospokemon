package connection

import (
	"encoding/json"
	"github.com/ospokemon/ospokemon/objects/spellscripts"
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
		action.Ability = walkAbility
	} else {
		action.Ability = entity.Controls().Abilities[ability]
	}

	log.Printf("Action accepted for client(%s)(%d): %v", client.Name, entityId, action)

	entity.Controls().Action = action
}

func copyMap(src map[string]*world.BasicView) map[string]*world.BasicView {
	dst := make(map[string]*world.BasicView)

	for k, v := range src {
		dst[k] = v
	}

	return dst
}

// Patch "Walk" ability onto every entity without occupying Ability slot

var walkAbility = &world.Ability{
	Spell: WalkSpell,
}

type walkSpell byte

var WalkSpell walkSpell
var walkCost = &world.SpellCost{0, make(map[int]int)}

func (walk walkSpell) Name() string {
	return "Walk"
}
func (walk walkSpell) Description() string {
	return "Walk"
}
func (walk walkSpell) CastTime() time.Duration {
	return 0
}
func (walk walkSpell) Cooldown() time.Duration {
	return 0
}
func (walk walkSpell) Cost() *world.SpellCost {
	return walkCost
}
func (walk walkSpell) Range() float64 {
	return 0
}
func (walk walkSpell) TargetType() world.TargetType {
	return world.TRGTnone
}
func (walk walkSpell) Script() world.SpellScript {
	return spellscripts.Scripts["Walk"]
}
