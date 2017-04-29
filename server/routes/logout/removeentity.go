package logout

import (
	"ospokemon.com"
)

func RemoveEntity(username string) {
	if player := ospokemon.Players.Cache[username]; player != nil {
		entity := player.GetEntity()
		universe := ospokemon.Multiverse[entity.UniverseId]

		universe.Remove(entity)
	}
}
