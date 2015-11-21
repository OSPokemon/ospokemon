package world

import (
	"time"
)

type EffectType uint8

const (
	EFCThealth EffectType = iota
	EFCTmove
	EFCTstun
	EFCTconjure
)

const (
	CTRLcollision    byte = 01
	CTRLfreeze       byte = 02
	CTRLstun         byte = 04
	CTRLroot         byte = 010
	CTRLinvulnerable byte = 020
)

type Effect struct {
	Name     string
	Type     EffectType
	Data     interface{}
	Start    time.Time
	Duration time.Duration
}

type Ability struct {
	Name        string
	Cast        time.Duration
	Cooldown    time.Duration
	MaxCooldown time.Duration
	Effects     []*Effect
}

var MoveAbility = &Ability{ // Special flag ability
	Name: "move",
}

type Action struct {
	Clock   time.Time
	Target  interface{}
	Ability *Ability
}

type Controls struct {
	Current   *Action
	State     byte
	Abilities map[string]*Ability
}

func (e *Effect) Update(entityId int, entity Entity, now time.Time) bool {
	var update bool

	switch e.Type {

	case EFCThealth:
		healthy := entity.(Healthy)
		power := e.Data.(int)
		newhealth := healthy.Health() + power

		if newhealth > healthy.MaxHealth() {
			power = healthy.MaxHealth() - healthy.Health()
		}
		if newhealth < 0 {
			power = 0 - healthy.Health()
		}

		if power != 0 {
			healthy.SetHealth(newhealth)
			update = true
		}

		break
	case EFCTmove:
		vector := e.Data.(*Vector)
		MoveEntity(entityId, vector)
		update = true
		break
	case EFCTstun:
		if int(entity.Controls().State&CTRLstun) > 0 {
			if e.Start.Add(e.Duration).Before(now) {
				entity.Controls().State ^= CTRLstun
				update = true
			}
		} else {
			entity.Graphics().Current = entity.Graphics().Animations[ANIMstun]
			entity.Controls().State |= CTRLstun
			entity.Controls().Current = nil
			update = true
		}
		break
	}

	return update
}

func (a *Action) Update(entityId int, entity Entity, now time.Time) bool {
	var update bool

	if a.Ability == MoveAbility {
		destination := a.Target.(Position)
		speedy := entity.(Speedy)
		vector := CreatePathVector(entity.Physics().Position, destination, speedy.Speed())
		moveEffect := &Effect{"move", EFCTmove, vector, now, 0}
		entity.SetEffects(append(entity.Effects(), moveEffect))
		update = true
	}

	return update
}
