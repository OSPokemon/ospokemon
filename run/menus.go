package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/save"
	"strings"
	"time"
)

const COMP_Menus = "menus"

type Menus struct {
	Player  bool
	Bag     bool
	Actions bool
}

func init() {
	event.On(event.PlayerMake, func(args ...interface{}) {
		p := args[0].(*save.Player)
		menus := &Menus{}

		p.Entity.AddComponent(menus)
	})

	event.On(event.BindingDown, func(args ...interface{}) {
		p := args[0].(*save.Player)
		b := args[1].(*save.Binding)

		if strings.HasPrefix(b.SystemId, "menu") {
			m := p.Entity.Component(COMP_Menus).(*Menus)
			m.Toggle(b.SystemId[5:])
		}
	})
}

func (m *Menus) Toggle(name string) {
	if name == "player" {
		m.Player = !m.Player
	} else if name == "bag" {
		m.Bag = !m.Bag
	} else if name == "actions" {
		m.Actions = !m.Actions
	}
}

func (m *Menus) Id() string {
	return COMP_Menus
}

func (m *Menus) Update(u *save.Universe, e *save.Entity, d time.Duration) {
}

func (m *Menus) Snapshot() map[string]interface{} {
	return nil
}
