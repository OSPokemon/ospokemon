package script

import (
	"github.com/ospokemon/ospokemon"
)

func MovementOn(player *ospokemon.Player, walk string) {
	movement := player.GetMovement()
	movement.Walk(walk)
}
