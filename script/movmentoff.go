package script

import (
	"ospokemon.com"
)

func MovementOff(player *ospokemon.Player, walk string) {
	movement := player.GetMovement()
	movement.ClearWalk(walk)
}
