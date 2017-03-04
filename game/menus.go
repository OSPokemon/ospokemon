package game

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/part"
)

type Menu string

type Menus map[Menu]bool

func MakeMenus() Menus {
	return map[Menu]bool{
		"chat":     false,
		"player":   false,
		"itembag":  false,
		"actions":  false,
		"settings": false,
	}
}

func (m Menus) Toggle(menu Menu) {
	if _, ok := m[menu]; !ok {
		logrus.WithFields(logrus.Fields{
			"name": menu,
		}).Warn("Menus: unrecognized menu name")
	}

	m[menu] = !m[menu]
}

func (m Menu) Part() string {
	return part.Menu
}

func (m Menus) Part() string {
	return part.Menus
}
