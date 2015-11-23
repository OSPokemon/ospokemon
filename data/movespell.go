package data

import (
	"github.com/ospokemon/api-go"
	"github.com/ospokemon/ospokemon/objects/spellscripts"
	"github.com/ospokemon/ospokemon/world"
	"time"
)

type MoveSpell struct {
	ospokemon.BasicMove
	CASTTIME float64
	COOLDOWN float64
	COST     *world.SpellCost
}

func (m *MoveSpell) CastTime() time.Duration {
	return time.Duration(m.CASTTIME)
}

func (m *MoveSpell) Cooldown() time.Duration {
	return time.Duration(m.COOLDOWN)
}

func (m *MoveSpell) Cost() *world.SpellCost {
	return m.COST
}

func (m *MoveSpell) Range() float64 {
	return m.Accuracy()
}

func (m *MoveSpell) Script() world.SpellScript {
	return spellscripts.Scripts[m.Name()]
}

func (m *MoveSpell) TargetType() world.TargetType {
	return world.TargetType(m.Targeter())
}
