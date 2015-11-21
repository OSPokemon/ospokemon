package world

import (
	"time"
)

type View struct {
	Name     string
	Physics  *Physics
	Graphic  string
	Controls ControlsView
	Effects  []EffectView
}

type ControlsView struct {
	Current   CastingView
	Abilities []AbilityView
}

type CastingView struct {
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

func MakeView(e Entity, now time.Time) *View {
	view := &View{}

	view.Name = e.Name()

	view.Physics = e.Physics()

	view.Graphic = e.Graphics().Current

	view.Controls = ControlsView{}
	view.Controls.Current = CastingView{}
	if e.Controls().Current != nil {
		view.Controls.Current.Name = e.Controls().Current.Ability.Name
		view.Controls.Current.Completion = float64(e.Controls().Current.Clock.Add(e.Controls().Current.Ability.Cast).Sub(now) / time.Minute)
	}
	view.Controls.Abilities = make([]AbilityView, len(e.Controls().Abilities))
	for _, ability := range e.Controls().Abilities {
		abilityView := AbilityView{
			Name:     ability.Name,
			Cooldown: float64(ability.Cooldown),
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
