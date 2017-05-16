package script

import (
	"ospokemon.com"
)

func MovementOn(player *ospokemon.Player, walk string) {
	movement := player.GetMovement()
	movement.Walk(walk)
}
