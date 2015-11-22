package update

import (
	"github.com/ospokemon/ospokemon/world"
	"time"
)

func UpdateEffect(effect *world.Effect, entity world.Entity, now time.Time) {
	isprotected := entity.Controls().State&world.CTRLPprotected < 1

	switch effect.Type {
	case world.EFCTimmune:
		entity.Controls().State |= world.CTRLimmune
		return
	case world.EFCTstasis:
		entity.Controls().State |= world.CTRLstasis
		return
	case world.EFCThealth:
		power := effect.Data.(int)
		if power < 0 && isprotected {
			return
		}

		healthy := entity.(world.Healthy)
		health := healthy.Health() + power

		if health > healthy.MaxHealth() {
			power = healthy.MaxHealth() - healthy.Health()
		} else if health < 0 {
			power = -healthy.Health()
		}

		healthy.SetHealth(healthy.Health() + power)
		return
	case world.EFCTstun:
		if isprotected {
			effect.Duration = 0
		}

		entity.Controls().State |= world.CTRLstun
		return
	case world.EFCTroot:
		if isprotected {
			effect.Duration = 0
		}

		entity.Controls().State |= world.CTRLroot
		return
	case world.EFCTmove:
		if entity.Controls().State&world.CTRLPstuck > 0 {
			return
		}

		vector := effect.Data.(*world.Vector)
		MoveEntity(entity, vector)
		return
	}
}
