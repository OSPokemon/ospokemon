package websocket

import (
	"encoding/json"

	"ospokemon.com"
	"ospokemon.com/script"
	"ospokemon.com/server/sessionman"
	"ztaylor.me/log"
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
	} else if message.Event == "MovementOn" {
		script.MovementOn(p, message.Message)
	} else if message.Event == "MovementOff" {
		script.MovementOff(p, message.Message)
	} else if message.Event == "MenuToggle" {
		p.GetMenus().Toggle(ospokemon.Menu(message.Message))
	} else if message.Event == "Chat" {
		script.SendChat(p, message.Message)
	} else if message.Event == "ItemCast" {
		itemcast(p, message.Message)
	} else if message.Event == "BindingSet" {
		bindingset(p, message.Message)
	} else if message.Event == "Click.Universe" {
		clickuniverse(p, message.Message)
	} else if message.Event == "Click.Entity" {
		clickentity(p, message.Message)
	} else if message.Event == "Dialog.Choice" {
		script.DialogChoice(p.GetEntity(), message.Message)
	} else {
		log.Add("Message", message).Add("Username", p.Username).Warn("receivemessage: unrecognized message type")
	}
}

func itemcast(player *ospokemon.Player, m string) {
	data := make(map[string]interface{})

	if err := json.Unmarshal([]byte(m), &data); err != nil {
		log.Add("Username", player.Username).Add("Error", err).Error("receivemessage: itemcast parse")
	} else if err := script.ItemCast(player.GetEntity(), data); err != nil {
		log.Add("Username", player.Username).Add("Error", err).Error("receivemessage: itemcast script")
	}
}

func bindingset(player *ospokemon.Player, m string) {
	data := make(map[string]interface{})

	if err := json.Unmarshal([]byte(m), &data); err != nil {
		log.Add("Username", player.Username).Add("Error", err).Error("receivemessage: bindingset parse")
	} else if err := script.BindingSet(player.GetEntity(), data); err != nil {
		log.Add("Username", player.Username).Add("Error", err).Error("receivemessage: bindingset script")
	}
}

func clickuniverse(player *ospokemon.Player, m string) {
	data := make(map[string]interface{})

	if err := json.Unmarshal([]byte(m), &data); err != nil {
		log.Add("Username", player.Username).Add("Error", err).Error("receivemessage: clickuniverse parse")
	} else if err := script.ClickUniverse(player.GetEntity(), data); err != nil {
		log.Add("Username", player.Username).Add("Error", err).Error("receivemessage: clickuniverse script")
	}
}

func clickentity(player *ospokemon.Player, m string) {
	data := make(map[string]interface{})

	if err := json.Unmarshal([]byte(m), &data); err != nil {
		log.Add("Username", player.Username).Add("Error", err).Error("receivemessage: clickuniverse parse")
	} else if err := script.ClickEntity(player.GetEntity(), data); err != nil {
		log.Add("Username", player.Username).Add("Error", err).Error("receivemessage: clickentity script")
	}
}
