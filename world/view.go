package world

import (
	"math"
	"time"
)

type BasicView struct {
	Id       int
	Name     string
	Physics  *Physics
	Graphic  string
	Portrait string
	Effects  []EffectView
}

type FullView struct {
	Id       int
	Name     string
	Physics  *Physics
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

func MakeBasicView(id int, e Entity, now time.Time) *BasicView {
	view := &BasicView{}

	view.Id = id

	view.Name = e.Name()

	view.Physics = e.Physics()

	view.Graphic = e.Graphics().Current

	view.Portrait = e.Graphics().Portrait

	view.Effects = make([]EffectView, 0)
	for _, effect := range e.Effects() {
		effectView := EffectView{
			Name:       effect.Name,
			Completion: float64(effect.Start.Add(effect.Duration).Sub(now) / time.Second),
		}
		view.Effects = append(view.Effects, effectView)
	}

	return view
}

func MakeFullView(id int, e Entity, now time.Time) *FullView {
	view := &FullView{}

	view.Id = id

	view.Name = e.Name()

	view.Physics = e.Physics()

	view.Graphic = e.Graphics().Current

	view.Portrait = e.Graphics().Portrait

	view.Controls = ControlsView{}
	view.Controls.Action = ActionView{}
	if e.Controls().Action != nil {
		view.Controls.Action.Name = e.Controls().Action.Ability.Spell.Name
		view.Controls.Action.Completion = float64(e.Controls().Action.Ability.LastCast.Add(e.Controls().Action.Ability.CastTime).Sub(now) / time.Second)
	}
	view.Controls.Abilities = make([]AbilityView, 0)
	for hotkey, ability := range e.Controls().Abilities {
		abilityView := AbilityView{
			Name:       ability.Spell.Name,
			Hotkey:     hotkey,
			TargetType: int(ability.Spell.TargetType),
			Graphic:    ability.Spell.Graphic,
			Cooldown:   math.Max(0, float64(ability.LastCast.Add(ability.Cooldown).Sub(now)/time.Second)),
		}
		view.Controls.Abilities = append(view.Controls.Abilities, abilityView)
	}

	view.Effects = make([]EffectView, 0)
	for _, effect := range e.Effects() {
		effectView := EffectView{
			Name:       effect.Name,
			Completion: float64(effect.Start.Add(effect.Duration).Sub(now) / time.Second),
		}
		view.Effects = append(view.Effects, effectView)
	}

	return view
}
