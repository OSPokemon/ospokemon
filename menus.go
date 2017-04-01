package ospokemon

import (
	"ospokemon.com/log"
)

const PARTmenu = "menu"
const PARTmenus = "menus"

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
		log.Add("name", menu).Warn("Menus: unrecognized menu name")
	}

	m[menu] = !m[menu]
}

func (m Menu) Part() string {
	return PARTmenu
}

func (parts Parts) GetMenu() Menu {
	menu, _ := parts[PARTmenu].(Menu)
	return menu
}

func (m Menus) Part() string {
	return PARTmenus
}

func (parts Parts) GetMenus() Menus {
	menus, _ := parts[PARTmenus].(Menus)
	return menus
}
