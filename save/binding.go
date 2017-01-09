package save

import (
	"github.com/ospokemon/ospokemon/part"
	"time"
)

type Binding struct {
	Key      string
	SystemId string
	part.Parts
}

func MakeBinding() *Binding {
	return &Binding{
		Parts: make(part.Parts),
	}
}

func (b *Binding) Update(u *Universe, e *Entity, d time.Duration) {
}

func (b *Binding) Json(expand bool) (string, map[string]interface{}) {
	data := map[string]interface{}{
		"key":      b.Key,
		"systemid": b.SystemId,
	}

	if expand {
		for _, part := range b.Parts {
			if jsoner, ok := part.(Jsoner); ok {
				key, partData := jsoner.Json(false)
				data[key] = partData
			}
		}
	}

	return "binding", data
}
