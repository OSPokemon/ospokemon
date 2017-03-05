package server

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/part"
	"github.com/ospokemon/ospokemon/space"
	// "github.com/ospokemon/ospokemon/run"
)

func ReceiveMessage(s *Session, m *WebsocketMessage) {
	p := game.Players[s.Username]

	s.Refresh()

	if m.Event == "Ping" {
		s.Send("Pong")
	} else if m.Event == "Key.Down" {
		keydown(p, m.Message)
	} else if m.Event == "Key.Up" {
		keyup(p, m.Message)
	} else if m.Event == "Binding.Set" {
		bindingset(p, m.Message)
	} else if m.Event == "Click.Universe" {
		clickuniverse(p, m.Message)
	} else if m.Event == "Click.Entity" {
		clickentity(p, m.Message)
	} else if m.Event == "Menu.Toggle" {
		menutoggle(p, m.Message)
	} else if m.Event == "Dialog.Choice" {
		dialogchoice(p, m.Message)
	} else {
		logrus.WithFields(logrus.Fields{
			"Message":  m,
			"Username": p.Username,
		}).Warn("Websocket: unrecognized message type")
	}
}

func keydown(p *game.Player, key string) {
	bindings := p.Parts[part.Bindings].(game.Bindings)
	binding := bindings[key]

	if binding == nil {
		return
	}

	event.Fire(event.BindingDown, p, binding)
}

func keyup(p *game.Player, key string) {
	bindings := p.Parts[part.Bindings].(game.Bindings)
	binding := bindings[key]

	if binding == nil {
		return
	}

	event.Fire(event.BindingUp, p, binding)
}

func bindingset(p *game.Player, m string) {
	data := make(map[string]interface{})
	json.Unmarshal([]byte(m), &data)

	key, ok := data["key"].(string)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"Username": p.Username,
			"Message":  m,
		}).Warn("Websocket: bindingset \"key\" missing")
		return
	}

	bindings := p.Parts[part.Bindings].(game.Bindings)

	binding := game.MakeBinding()
	binding.Key = key

	if data["spell"] != nil {
		spellId := uint(data["spell"].(float64))
		actions := p.Parts[part.Actions].(game.Actions)
		action := actions[spellId]

		if action == nil {
			return
		}

		if action.Parts[part.Bindings] == nil {
			action.AddPart(make(game.Bindings))
		}
		action.Parts[part.Bindings].(game.Bindings)[binding.Key] = binding

		binding.AddPart(action.Parts[part.Imaging])

	} else if data["itemslot"] != nil {
		itemslotid := uint(data["itemslot"].(float64))
		itembag := p.Parts[part.Itembag].(*game.Itembag)
		itemslot := itembag.Slots[itemslotid]

		if itemslot == nil {
			return
		}
		if oldBinding, _ := itemslot.Parts[part.Binding].(*game.Binding); oldBinding != nil {
			oldBinding.RemovePart(oldBinding)
			delete(bindings, oldBinding.Key)
		}

		itemslot.AddPart(binding)
		binding.Parts = itemslot.Parts
	} else if data["walk"] != nil {
		direction := data["walk"].(string)

		for k, b := range bindings {
			if b.Parts[part.Walk] != nil {
				if direction == string(b.Parts[part.Walk].(game.Walk)) {
					delete(bindings, k)
				}
			}
		}

		binding.AddPart(game.Walk(direction))
		imaging := game.MakeImaging()
		imaging.Image = "/img/ui/walk/" + direction + ".png"
		binding.AddPart(imaging)
	} else if data["menu"] != nil {
		menu := data["menu"].(string)

		for k, b := range bindings {
			if b.Parts[part.Menu] != nil {
				if menu == string(b.Parts[part.Menu].(game.Menu)) {
					delete(bindings, k)
				}
			}
		}

		binding.AddPart(game.Menu(menu))
		imaging := game.MakeImaging()
		imaging.Image = "/img/ui/menu/" + menu + ".png"
		binding.AddPart(imaging)
	} else {
		logrus.WithFields(logrus.Fields{
			"Message":  m,
			"Username": p.Username,
		}).Warn("Websocket: unrecognized Binding.Set message")
	}

	bindings[binding.Key] = binding
}

func clickuniverse(player *game.Player, m string) {
	data := make(map[string]interface{})
	json.Unmarshal([]byte(m), &data)

	point := space.Point{}
	if pointx, ok := data["x"].(float64); ok {
		point.X = pointx
	} else {
		return
	}
	if pointy, ok := data["y"].(float64); ok {
		point.Y = pointy
	} else {
		return
	}

	movement := player.Parts[part.Movement].(*game.Movement)
	movement.Up = false
	movement.Down = false
	movement.Left = false
	movement.Right = false
	movement.Target = &point
}

func clickentity(player *game.Player, m string) {
	data := make(map[string]interface{})
	json.Unmarshal([]byte(m), &data)

	if entityId, ok := data["entity"].(float64); ok {
		playerEntity := player.Parts[part.Entity].(*game.Entity)
		universe := game.Multiverse[playerEntity.UniverseId]
		entity := universe.Entities[uint(entityId)]

		if dialog, _ := entity.Parts[part.Dialog].(*game.Dialog); dialog != nil {
			player.AddPart(dialog)
		}
	}
}

func dialogchoice(player *game.Player, m string) {
	entity := player.Parts[part.Entity].(*game.Entity)

	if dialog, ok := player.Parts[part.Dialog].(*game.Dialog); ok {
		if nextDialog := dialog.Next(m); nextDialog != nil {

			if script := game.Scripts[nextDialog.Script]; script != nil {
				if err := script(entity, nextDialog.Data); err != nil {
					logrus.WithFields(logrus.Fields{
						"Player": player.Username,
						"Choice": m,
						"Error":  err.Error(),
					}).Error("dialog choice")
					return
				}
			}

			player.AddPart(nextDialog)
		} else if len(dialog.Choices) < 1 {
			player.RemovePart(dialog)
		}
	}
}

func menutoggle(p *game.Player, m string) {
	menus := p.Parts[part.Menus].(game.Menus)
	menus.Toggle(game.Menu(m))
}
