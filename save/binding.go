package save

import (
	"time"
)

type Binding struct {
	Key      string
	SpellId  uint
	BagSlot  int
	SystemId string
	Image    string
	Timer    *time.Duration
}

func MakeBinding(key string) *Binding {
	return &Binding{
		Key:     key,
		BagSlot: -1,
	}
}

func (b *Binding) Update(u *Universe, e *Entity, d time.Duration) {
}

func (b *Binding) Snapshot() map[string]interface{} {
	timebuff := 0
	if b.Timer != nil {
		timebuff = int(*b.Timer)
	}

	return map[string]interface{}{
		"key":      b.Key,
		"image":    b.Image,
		"spellid":  b.SpellId,
		"bagslot":  b.BagSlot,
		"systemid": b.SystemId,
		"timer":    timebuff,
	}
}
