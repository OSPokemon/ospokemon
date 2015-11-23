package world

import (
	"time"
)

type BasicView struct {
	Id      int
	Name    string
	Physics *Physics
	Graphic string
	Effects []EffectView
}

type FullView struct {
	Id       int
	Name     string
	Physics  *Physics
	Graphic  string
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
	Name     string
	Cooldown float64
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

	view.Effects = make([]EffectView, len(e.Effects()))
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

	view.Controls = ControlsView{}
	view.Controls.Action = ActionView{}
	if e.Controls().Action != nil {
		view.Controls.Action.Name = e.Controls().Action.Ability.Spell.Name()
		view.Controls.Action.Completion = float64(e.Controls().Action.Clock.Add(e.Controls().Action.Ability.Spell.CastTime()).Sub(now) / time.Second)
	}
	view.Controls.Abilities = make([]AbilityView, len(e.Controls().Abilities))
	for _, ability := range e.Controls().Abilities {
		abilityView := AbilityView{
			Name:     ability.Spell.Name(),
			Cooldown: float64(ability.LastCast.Add(ability.Spell.Cooldown()).Sub(now) / time.Second),
		}
		view.Controls.Abilities = append(view.Controls.Abilities, abilityView)
	}

	view.Effects = make([]EffectView, len(e.Effects()))
	for _, effect := range e.Effects() {
		effectView := EffectView{
			Name:       effect.Name,
			Completion: float64(effect.Start.Add(effect.Duration).Sub(now) / time.Second),
		}
		view.Effects = append(view.Effects, effectView)
	}

	return view
}
