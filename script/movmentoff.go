package script

import (
	"github.com/ospokemon/ospokemon"
)

func MovementOff(player *ospokemon.Player, walk string) {
	movement := player.GetMovement()
	movement.ClearWalk(walk)
}
