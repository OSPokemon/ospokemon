package engine

import (
	"time"
)

func UpdateMap(m *Map, now time.Time) {
	m.Lock()
	defer m.Unlock()

	t := now.Sub(m.LastUpdate)

	for _, entityId := range m.Entities {
		e := Entities[entityId]
		UpdateEntity(m, e, now, t)
	}

	if m.MapScript != nil {
		go m.MapScript(m, now, t)
	}
}
