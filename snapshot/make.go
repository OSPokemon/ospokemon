package snapshot

import (
	"github.com/ospokemon/ospokemon/engine"
	"math"
	"strconv"
	"time"
)

func Make(m *engine.Map, now time.Time) (map[string]interface{}, map[string]interface{}) {
	view := make(map[string]interface{})
	cview := make(map[string]interface{})

	for _, entityId := range m.Entities {
		entity := engine.Entities[entityId]
		view[strconv.Itoa(entityId)] = MakeBasicView(entityId, entity, now)
		cview[strconv.Itoa(entityId)] = MakeFullView(entityId, entity, now)
	}

	return view, cview
}

func MakeBasicView(id int, e engine.Entity, now time.Time) map[string]interface{} {
	view := make(map[string]interface{})

	view["Id"] = id
	view["Graphic"] = e.Graphic()
	view["Shape"] = e.Shape()

	entity, ok := e.(engine.LivingEntity)
	if !ok {
		return view
	}

	view["Name"] = entity.Name()
	view["Portrait"] = entity.Graphics()[engine.ANIMportrait]

	view["Action"] = MakeActionView(entity.Action(), now)

	view["Status"] = MakeControlDescriptors(*entity.Control())

	view["Effects"] = MakeEffectsDescriptors(*entity.Effects())

	return view
}

func MakeActionView(action *engine.Action, now time.Time) interface{} {
	if action == nil {
		return false
	}

	view := make(map[string]interface{})

	view["Name"] = action.Ability.Spell.Name
	view["Graphic"] = action.Ability.Spell.Graphic

	if action.ChannelStart != nil {
		view["Completion"] = 100 - float64(now.Sub(*action.ChannelStart))/float64(action.Ability.ChannelTime)*100
	} else if action.CastStart != nil {
		view["Completion"] = float64(now.Sub(*action.CastStart)) / float64(action.Ability.CastTime) * 100
	} else {

	}

	return view
}

func MakeControlDescriptors(control engine.Control) []string {
	status := make([]string, 0)
	if engine.CTRLdead&control > 0 {
		status = append(status, "dead")
	}
	if engine.CTRLimmune&control > 0 {
		status = append(status, "immune")
	}
	if engine.CTRLstasis&control > 0 {
		status = append(status, "stasis")
	}
	if engine.CTRLstun&control > 0 {
		status = append(status, "stun")
	}
	if engine.CTRLroot&control > 0 {
		status = append(status, "root")
	}
	if engine.CTRLcloak&control > 0 {
		status = append(status, "cloak")
	}
	if engine.CTRLsilence&control > 0 {
		status = append(status, "silence")
	}
	return status
}

func MakeEffectsDescriptors(effects []*engine.Effect) []string {
	fx := make([]string, 0)
	for _, effect := range effects {
		fx = append(fx, effect.Name)
	}
	return fx
}

func MakeFullView(id int, e engine.Entity, now time.Time) map[string]interface{} {
	view := MakeBasicView(id, e, now)

	entity, ok := e.(engine.LivingEntity)
	if !ok {
		return view
	}

	abilities := make([]map[string]interface{}, 0)

	for hotkey, ability := range *entity.Abilities() {
		abilityview := make(map[string]interface{})
		abilityview["Name"] = ability.Spell.Name
		abilityview["Hotkey"] = hotkey
		abilityview["TargetType"] = ability.Spell.TargetType
		abilityview["Graphic"] = ability.Spell.Graphic

		abilityview["Cooldown"] = roundTime(ability.LastCast.Add(ability.Cooldown).Sub(now))

		abilities = append(abilities, abilityview)
	}

	view["Abilities"] = abilities

	return view
}

func roundTime(duration time.Duration) float64 {
	timeinitial := float64(duration / time.Millisecond / 100)
	time := math.Floor(timeinitial)
	time = math.Max(0, time/10)

	return time
}
