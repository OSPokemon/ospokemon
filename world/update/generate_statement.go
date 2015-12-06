package update

import (
	"github.com/ospokemon/ospokemon/world"
	"math"
	"time"
)

type BasicView struct {
	Id       int
	Name     string
	Physics  *world.Physics
	Graphic  string
	Portrait string
	Effects  []EffectView
}

type FullView struct {
	Id       int
	Name     string
	Physics  *world.Physics
	Graphic  string
	Portrait string
	Controls ControlsView
	Effects  []EffectView
}

type ControlsView struct {
	Action    ActionView
	Abilities []AbilityView
}

type ActionView struct {
	Name       string
	Completion float64
}

type AbilityView struct {
	Name       string
	Hotkey     string
	TargetType int
	Graphic    string
	Cooldown   float64
}

type EffectView struct {
	Name       string
	Completion float64
}

func MakeBasicView(id int, e world.Entity, now time.Time) map[string]interface{} {
	view := make(map[string]interface{})

	view["Id"] = id
	view["Name"] = e.Name()
	view["Portrait"] = e.Graphics().Portrait
	view["Graphic"] = e.Graphics().Current
	view["Physics"] = e.Physics()

	if e.Action() != nil && e.Action().Start != nil {
		action := make(map[string]interface{})
		action["Name"] = e.Action().Ability.Spell.Name
		action["Completion"] = float64(e.Action().Start.Add(e.Action().Ability.CastTime).Sub(now))
		action["Graphic"] = e.Action().Ability.Spell.Graphic

		view["Action"] = action
	}
	if mortal, ok := e.(world.Mortality); ok {
		effects := make([]map[string]interface{}, 0)
		for _, effect := range mortal.Effects() {
			if effect.Start == nil {
				continue
			}

			effectView := map[string]interface{}{
				"Name":       effect.Name,
				"Completion": float64(effect.Start.Add(effect.Duration).Sub(now) / time.Second),
			}
			effects = append(effects, effectView)
		}
		view["Effects"] = effects
	}

	return view
}

func MakeFullView(id int, e world.Entity, now time.Time) map[string]interface{} {
	view := MakeBasicView(id, e, now)

	// Extra stuff
	if intelligent, ok := e.(world.Intelligence); ok {
		abilities := make([]map[string]interface{}, 0)

		for hotkey, ability := range intelligent.Abilities() {
			abilityview := make(map[string]interface{})
			abilityview["Name"] = ability.Spell.Name
			abilityview["Hotkey"] = hotkey
			abilityview["TargetType"] = int(ability.Spell.TargetType)
			abilityview["Graphic"] = ability.Spell.Graphic

			abilityview["Cooldown"] = math.Max(0, float64(ability.LastCast.Add(ability.Cooldown).Sub(now)/time.Second))

			abilities = append(abilities, abilityview)
		}

		view["Abilities"] = abilities
	}

	return view
}
