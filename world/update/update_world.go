package update

import (
	"github.com/ospokemon/ospokemon/world"
	"strconv"
	"time"
)

func UpdateWorld(now time.Time) map[string]interface{} {
	view := make(map[string]interface{})

	for _, entity := range world.Entities {
		UpdateEntity(entity, now)
	}

	for id, entity := range world.Entities {
		eview := MakeBasicView(id, entity, now)
		view[strconv.Itoa(id)] = eview
	}

	return view
}
