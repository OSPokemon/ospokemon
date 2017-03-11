package game

import (
	"github.com/ospokemon/ospokemon/part"
	"time"
)

type ChatMessage struct {
	Message string
	Timer   *time.Duration
}

func (m *ChatMessage) Part() string {
	return part.ChatMessage
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
