package run

import (
	"ospokemon.com/event"
	"ospokemon.com/game"
	"ospokemon.com/log"
)

func init() {
	event.On(event.ActionCast, ActionCast)
}

func ActionCast(args ...interface{}) {
	universe := args[0].(*game.Universe)
	entity := args[1].(*game.Entity)
	action := args[2].(*game.Action)

	spell := action.Spell
	timer := spell.Cooldown
	action.Timer = &timer

	if script, ok := game.Scripts[spell.Script]; ok {
		if err := script(entity, spell.Data); err != nil {
			log.Add("Universe", universe.Id).Add("Entity", entity.Id).Add("Spell", spell.Id).Add("Error", err.Error()).Error("action cast")
		}
	}
}
