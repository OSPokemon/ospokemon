package run

import (
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/part"
	"github.com/ospokemon/ospokemon/save"
	"strings"
)

type Menus struct {
	Player  bool
	Itembag bool
	Actions bool
}

func init() {
	event.On(event.PlayerMake, func(args ...interface{}) {
		p := args[0].(*save.Player)
		menus := &Menus{}

		p.AddPart(menus)
	})

	event.On(event.BindingDown, func(args ...interface{}) {
		p := args[0].(*save.Player)
		b := args[1].(*save.Binding)

		if strings.HasPrefix(b.SystemId, "menu") {
			m := p.Parts[part.MENUS].(*Menus)
			m.Toggle(b.SystemId[5:])
		}
	})
}

func (m *Menus) Toggle(name string) {
	if name == "player" {
		m.Player = !m.Player
	} else if name == "itembag" {
		m.Itembag = !m.Itembag
	} else if name == "actions" {
		m.Actions = !m.Actions
	}
}

func (m *Menus) Part() string {
	return part.MENUS
}
