package update

import (
	"github.com/ospokemon/ospokemon/world"
	"strconv"
	"time"
)

func UpdateWorld(now time.Time) map[string]*world.View {
	view := make(map[string]*world.View)

	for _, entity := range world.Entities {
		UpdateEntity(entity, now)
	}

	for id, entity := range world.Entities {
		eview := world.MakeView(id, entity, now)
		view[strconv.Itoa(id)] = eview
	}

	return view
}
