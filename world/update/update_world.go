package update

import (
	"github.com/ospokemon/ospokemon/world"
	"strconv"
	"time"
)

func UpdateWorld(now time.Time) map[string]*world.BasicView {
	view := make(map[string]*world.BasicView)

	for _, entity := range world.Entities {
		ResetEntity(entity)
		UpdateEntity(entity, now)
	}

	for id, entity := range world.Entities {
		if entity.Controls().State&world.CTRLcloak > 0 {
			continue
		}

		eview := world.MakeBasicView(id, entity, now)
		view[strconv.Itoa(id)] = eview
	}

	return view
}