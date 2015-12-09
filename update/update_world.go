package update

import (
	"github.com/ospokemon/ospokemon/world"
	"strconv"
	"time"
)

func UpdateWorld(now time.Time) (map[string]interface{}, map[string]interface{}) {
	view := make(map[string]interface{})
	cview := make(map[string]interface{})

	for _, entity := range world.Entities {
		UpdateEntity(entity, now)
	}

	for id, entity := range world.Entities {
		view[strconv.Itoa(id)] = MakeBasicView(id, entity, now)
		cview[strconv.Itoa(id)] = MakeFullView(id, entity, now)
	}

	return view, cview
}
