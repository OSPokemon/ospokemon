package server

import (
	"encoding/json"
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"ospokemon.com/script"
	"ospokemon.com/server/session"
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

func itemcast(player *ospokemon.Player, m string) {
	data := make(map[string]interface{})

	if err := json.Unmarshal([]byte(m), &data); err != nil {
		log.Add("Error", err).Error("receivemessage: itemcast")
	} else if err := script.ItemCast(player.GetEntity(), data); err != nil {
		log.Add("Username", player.Username).Add("Error", err).Error("receivemessage: itemcast")
	}
}

func bindingset(player *ospokemon.Player, m string) {
	data := make(map[string]interface{})

	if err := json.Unmarshal([]byte(m), &data); err != nil {
		log.Add("Error", err).Error("receivemessage: bindingset")
	} else if err := script.BindingSet(player.GetEntity(), data); err != nil {
		log.Add("Username", player.Username).Add("Error", err).Error("receivemessage: bindingset")
	}
}

func clickuniverse(player *ospokemon.Player, m string) {
	data := make(map[string]interface{})

	if err := json.Unmarshal([]byte(m), &data); err != nil {
		log.Add("Error", err).Error("receivemessage: clickuniverse")
	} else if err := script.ClickUniverse(player.GetEntity(), data); err != nil {
		log.Add("Username", player.Username).Add("Error", err).Error("receivemessage: clickuniverse")
	}
}

func clickentity(player *ospokemon.Player, m string) {
	data := make(map[string]interface{})

	if err := json.Unmarshal([]byte(m), &data); err != nil {
		log.Add("Error", err).Error("receivemessage: clickuniverse")
	} else if err := script.ClickEntity(player.GetEntity(), data); err != nil {
		log.Add("Username", player.Username).Add("Error", err).Error("receivemessage: clickentity")
	}
}

func dialogchoice(player *ospokemon.Player, m string) {
	data := make(map[string]interface{})

	if err := json.Unmarshal([]byte(m), &data); err != nil {
		log.Add("Error", err).Error("receivemessage: dialogchoice")
	} else if err := script.DialogChoice(player.GetEntity(), data); err != nil {
		log.Add("Username", player.Username).Add("Error", err.Error()).Error("receivemessage: itemcast")
	}
}

func menutoggle(player *ospokemon.Player, m string) {
	player.GetMenus().Toggle(ospokemon.Menu(m))
}

func chat(player *ospokemon.Player, m string) {
	timer := 3 * time.Second
	chatmessage := &ospokemon.ChatMessage{
		Message: m,
		Timer:   &timer,
	}
	player.AddPart(chatmessage)
}
