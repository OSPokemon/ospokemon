package server

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/part"
	"github.com/ospokemon/ospokemon/script"
	"github.com/ospokemon/ospokemon/space"
	"time"
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
	} else if m.Event == "Chat" {
		chat(p, m.Message)
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
	err := script.BindingSet(p.Parts[part.Entity].(*game.Entity), data)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Username": p.Username,
		}).Error(err.Error())
	}
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

func chat(p *game.Player, m string) {
	timer := 3 * time.Second
	chatmessage := &game.ChatMessage{
		Message: m,
		Timer:   &timer,
	}
	p.AddPart(chatmessage)
}
