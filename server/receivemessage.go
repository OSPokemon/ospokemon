package server

import (
	"encoding/json"
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"ospokemon.com/script"
	"ospokemon.com/server/session"
	"ospokemon.com/space"
	"strconv"
	"time"
)

func ReceiveMessage(s *session.Session, m *session.WebsocketMessage) {
	p, _ := ospokemon.GetPlayer(s.Username)

	s.Refresh()

	if m.Event == "Ping" {
		p.GetToaster().Add(&ospokemon.Toast{
			Color:   "blue",
			Message: "Pong",
			Image:   "/img/ospokemon.png",
		})
		s.Send("Pong")
	} else if m.Event == "Key.Down" {
		keydown(p, m.Message)
	} else if m.Event == "Key.Up" {
		keyup(p, m.Message)
	} else if m.Event == "Item.Cast" {
		itemcast(p, m.Message)
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

func keydown(player *ospokemon.Player, key string) {
	if binding := player.GetBindings()[key]; binding != nil {
		event.Fire(event.BindingDown, player, binding)
	}
}

func keyup(player *ospokemon.Player, key string) {
	if binding := player.GetBindings()[key]; binding != nil {
		event.Fire(event.BindingUp, player, binding)
	}
}

func itemcast(player *ospokemon.Player, itemidS string) {
	itemidI, err := strconv.Atoi(itemidS)
	if err != nil {
		log.Add("Error", err).Error("receivemessage: itemcast")
		return
	}

	itemid := uint(itemidI)

	itembag := player.GetItembag()
	if itembag.Timers[itemid] != nil {
		log.Add("Player", player.Username).Add("ItemId", itemid).Add("Timer", itembag.Timers[itemid]).Debug("receivemessage: itemcast: item cooldown")
		return
	}

	itemslot := itembag.Slots[itemid]
	if itemslot == nil {
		log.Add("Error", err).Add("ItemId", itemid).Error("receivemessage: item not found")
		return
	}

	item := itemslot.Item
	timer := item.CastTime + item.Cooldown
	itembag.Timers[itemslot.Item.Id] = &timer
}

func bindingset(player *ospokemon.Player, m string) {
	data := make(map[string]interface{})
	json.Unmarshal([]byte(m), &data)

	if err := script.BindingSet(player.GetEntity(), data); err != nil {
		log.Add("Username", player.Username).Error(err.Error())
	}
}

func clickuniverse(player *ospokemon.Player, m string) {
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

func clickentity(player *ospokemon.Player, m string) {
	data := make(map[string]interface{})
	json.Unmarshal([]byte(m), &data)

	if entityId, ok := data["entity"].(float64); ok {
		playerEntity := player.GetEntity()
		universe := ospokemon.Multiverse[playerEntity.UniverseId]
		entity := universe.Entities[uint(entityId)]

		if dialog := entity.GetDialog(); dialog != nil {
			player.AddPart(dialog)
		}
	}
}

func dialogchoice(player *ospokemon.Player, m string) {
	entity := player.GetEntity()

	if dialog := player.GetDialog(); dialog != nil {
		if nextDialog := dialog.Next(m); nextDialog != nil {
			for _, tester := range nextDialog.Tests {
				if !tester.Test(player) {
					return
				}
			}

			player.AddPart(nextDialog)

			for _, scripter := range nextDialog.Scripts {
				if err := scripter.Run(entity); err != nil {
					log.Add("Player", player.Username).Add("Choice", m).Add("Error", err.Error()).Error("dialog choice")
					return
				}
			}
		} else if len(dialog.Choices) < 1 {
			player.RemovePart(dialog)
		}
	}
}

func menutoggle(player *ospokemon.Player, m string) {
	menus := player.GetMenus()
	menus.Toggle(ospokemon.Menu(m))
}

func chat(player *ospokemon.Player, m string) {
	timer := 3 * time.Second
	chatmessage := &ospokemon.ChatMessage{
		Message: m,
		Timer:   &timer,
	}
	player.AddPart(chatmessage)
}
