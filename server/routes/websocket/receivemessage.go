package websocket

import (
	"encoding/json"
	"ospokemon.com"
	"ospokemon.com/event"
	"ospokemon.com/log"
	"ospokemon.com/script"
	"ospokemon.com/server/sessionman"
	"time"
)

func ReceiveMessage(session *sessionman.Session, message *sessionman.WebsocketMessage) {
	p, _ := ospokemon.GetPlayer(session.Username)

	session.Refresh()

	if message.Event == "Ping" {
		p.GetToaster().Add(&ospokemon.Toast{
			Color:   "blue",
			Message: "Pong",
			Image:   "/img/ospokemon.png",
		})
		session.Send("Pong")
	} else if message.Event == "Key.Down" {
		keydown(p, message.Message)
	} else if message.Event == "Key.Up" {
		keyup(p, message.Message)
	} else if message.Event == "Item.Cast" {
		itemcast(p, message.Message)
	} else if message.Event == "Binding.Set" {
		bindingset(p, message.Message)
	} else if message.Event == "Click.Universe" {
		clickuniverse(p, message.Message)
	} else if message.Event == "Click.Entity" {
		clickentity(p, message.Message)
	} else if message.Event == "Menu.Toggle" {
		menutoggle(p, message.Message)
	} else if message.Event == "Dialog.Choice" {
		dialogchoice(p, message.Message)
	} else if message.Event == "Chat" {
		chat(p, message.Message)
	} else {
		log.Add("Message", message).Add("Username", p.Username).Warn("ReceiveMessage: unrecognized message type")
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
