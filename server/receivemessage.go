package server

import (
	"encoding/json"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/log"
	"github.com/ospokemon/ospokemon/script"
	"github.com/ospokemon/ospokemon/space"
	"time"
)

func ReceiveMessage(s *Session, m *WebsocketMessage) {
	p := game.Players[s.Username]

	s.Refresh()

	if m.Event == "Ping" {
		p.GetToaster().Add(&game.Toast{
			Color:   "blue",
			Message: "Pong",
			Image:   "/img/ospokemon.png",
		})
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
		log.Add("Message", m).Add("Username", p.Username).Warn("ReceiveMessage: unrecognized message type")
	}
}

func keydown(player *game.Player, key string) {
	if binding := player.GetBindings()[key]; binding != nil {
		event.Fire(event.BindingDown, player, binding)
	}
}

func keyup(player *game.Player, key string) {
	if binding := player.GetBindings()[key]; binding != nil {
		event.Fire(event.BindingUp, player, binding)
	}
}

func bindingset(player *game.Player, m string) {
	data := make(map[string]interface{})
	json.Unmarshal([]byte(m), &data)

	if err := script.BindingSet(player.GetEntity(), data); err != nil {
		log.Add("Username", player.Username).Error(err.Error())
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

	movement := player.GetMovement()
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
		playerEntity := player.GetEntity()
		universe := game.Multiverse[playerEntity.UniverseId]
		entity := universe.Entities[uint(entityId)]

		if dialog := entity.GetDialog(); dialog != nil {
			player.AddPart(dialog)
		}
	}
}

func dialogchoice(player *game.Player, m string) {
	entity := player.GetEntity()

	if dialog := player.GetDialog(); dialog != nil {
		if nextDialog := dialog.Next(m); nextDialog != nil {

			if script := game.Scripts[nextDialog.Script]; script != nil {
				if err := script(entity, nextDialog.Data); err != nil {
					log.Add("Player", player.Username).Add("Choice", m).Add("Error", err.Error()).Error("dialog choice")
					return
				}
			}

			player.AddPart(nextDialog)
		} else if len(dialog.Choices) < 1 {
			player.RemovePart(dialog)
		}
	}
}

func menutoggle(player *game.Player, m string) {
	menus := player.GetMenus()
	menus.Toggle(game.Menu(m))
}

func chat(player *game.Player, m string) {
	timer := 3 * time.Second
	chatmessage := &game.ChatMessage{
		Message: m,
		Timer:   &timer,
	}
	player.AddPart(chatmessage)
}
