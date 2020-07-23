package script

import (
	"github.com/ospokemon/ospokemon"
	"time"
)

func SendChat(player *ospokemon.Player, m string) {
	timer := 3 * time.Second
	chatmessage := &ospokemon.ChatMessage{
		Message: m,
		Timer:   &timer,
	}
	player.AddPart(chatmessage)
}
