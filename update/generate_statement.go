package update

import (
	"github.com/ospokemon/ospokemon/world"
	"math"
	"time"
)

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
		action["Completion"] = float64(now.Sub(*e.Action().Start)) / float64(e.Action().Ability.CastTime) * 100
		action["Graphic"] = e.Action().Ability.Spell.Graphic

		view["Action"] = action
	} else {
		view["Action"] = false
	}

	if mortal, ok := e.(world.Mortality); ok {
		effects := make([]map[string]interface{}, 0)
		for _, effect := range mortal.Effects() {
			if effect.Start == nil {
				continue
			}

			effectView := map[string]interface{}{
				"Name":       effect.Name,
				"Completion": roundTime(effect.Start.Add(effect.Duration).Sub(now)),
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

			abilityview["Cooldown"] = roundTime(ability.LastCast.Add(ability.Cooldown).Sub(now))

			abilities = append(abilities, abilityview)
		}

		view["Abilities"] = abilities
	}

	return view
}

func roundTime(duration time.Duration) float64 {
	timeinitial := float64(duration / time.Millisecond / 100)
	time := math.Floor(timeinitial)
	time = math.Max(0, time/10)

	return time
}
