package spellscripts

import (
	"github.com/ospokemon/ospokemon/world"
	"log"
	"time"
)

func init() {
	Scripts["Walk"] = Walk
}

func Walk(self world.Entity, t interface{}, now time.Time) {
	position := self.Physics().Position
	destination := t.(*world.Position)

	if world.GetDistance(&position, destination) < 10 {
		self.Controls().Action = nil
		return
	}

	speedy := self.(world.Speedy)
	vector := world.CreatePathVector(&self.Physics().Position, destination, speedy.Speed())

	moveEffect := &world.Effect{"walk", world.EFCTmove, vector, now, 0}
	self.SetEffects(append(self.Effects(), moveEffect))

	log.Printf("MoveAction %s along %v towards %v", self.Name(), vector, destination)
}
