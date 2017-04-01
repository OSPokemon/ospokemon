package ospokemon

import (
	"time"
)

const PARTchatmessage = "chat"

type ChatMessage struct {
	Message string
	Timer   *time.Duration
}

func (m *ChatMessage) Part() string {
	return PARTchatmessage
}

func (parts Parts) GetChatMessage() *ChatMessage {
	chat, _ := parts[PARTchatmessage].(*ChatMessage)
	return chat
}

func (m *ChatMessage) Update(u *Universe, e *Entity, t time.Duration) {
	if m.Timer == nil {
		e.RemovePart(m)
		return
	}

	*m.Timer -= t
	if *m.Timer < 0 {
		m.Timer = nil
	}
}
