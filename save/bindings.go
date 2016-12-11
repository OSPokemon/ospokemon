package save

import (
	"github.com/ospokemon/ospokemon/event"
	"time"
)

const COMP_Bindings = "bindings"

type Bindings map[string]*Binding

func init() {
	event.On(event.PlayerMake, func(args ...interface{}) {
		p := args[0].(*Player)
		bindings := make(Bindings)

		p.Entity.AddComponent(bindings)
	})

	event.On(event.PlayerQuery, func(args ...interface{}) {
		p := args[0].(*Player)
		bindings := p.Entity.Component(COMP_Bindings).(Bindings)
		bindings.QueryPlayer(p.Username)

		if len(bindings) == 0 {
			for key, binding := range MakeBindingsDefault() {
				bindings[key] = binding
			}
		}
	})

	event.On(event.PlayerInsert, func(args ...interface{}) {
		p := args[0].(*Player)
		bindings := p.Entity.Component(COMP_Bindings).(Bindings)
		bindings.InsertPlayer(p.Username)
	})

	event.On(event.PlayerDelete, func(args ...interface{}) {
		p := args[0].(*Player)
		bindings := p.Entity.Component(COMP_Bindings).(Bindings)
		bindings.DeletePlayer(p.Username)
	})
}

func MakeBindingsDefault() Bindings {
	return Bindings{
		"w": &Binding{Key: "w", BagSlot: -1, SystemId: "walk-up"},
		"d": &Binding{Key: "d", BagSlot: -1, SystemId: "walk-right"},
		"a": &Binding{Key: "a", BagSlot: -1, SystemId: "walk-left"},
		"s": &Binding{Key: "s", BagSlot: -1, SystemId: "walk-down"},
		"c": &Binding{Key: "c", BagSlot: -1, SystemId: "menu-player", Image: "/img/ui/player.png"},
		"b": &Binding{Key: "b", BagSlot: -1, SystemId: "menu-bag", Image: "/img/ui/bag.png"},
		"v": &Binding{Key: "v", BagSlot: -1, SystemId: "menu-actions", Image: "/img/ui/actions.png"},
		"z": &Binding{Key: "z", BagSlot: -1, SystemId: "menu-bindings"},
	}
}

func (b Bindings) clear() {
	for k := range b {
		delete(b, k)
	}
}

func (b Bindings) Id() string {
	return COMP_Bindings
}

func (b Bindings) Update(u *Universe, e *Entity, d time.Duration) {
	for _, binding := range b {
		binding.Update(u, e, d)
	}
}

func (b Bindings) Snapshot() map[string]interface{} {
	return nil
}

func (b Bindings) SnapshotDetail() map[string]interface{} {
	data := make(map[string]interface{})

	for key, binding := range b {
		bdata := binding.Snapshot()
		bdata["key"] = key
		data[key] = bdata
	}

	return data
}

func (b Bindings) QueryPlayer(username string) error {
	rows, err := Connection.Query(
		"SELECT key, spellid, bagslot, systemid, image FROM bindings_players WHERE username=?",
		username,
	)

	if err != nil {
		return err
	}

	b.clear()
	for rows.Next() {
		binding := &Binding{}

		err = rows.Scan(&binding.Key, &binding.SpellId, &binding.BagSlot, &binding.SystemId, &binding.Image)
		if err != nil {
			return err
		}

		b[binding.Key] = binding
	}
	rows.Close()

	return nil
}

func (b Bindings) InsertPlayer(username string) error {
	for key, binding := range b {
		_, err := Connection.Exec(
			"INSERT INTO bindings_players (username, key, spellid, bagslot, systemid, image) VALUES (?, ?, ?)",
			username,
			key,
			binding.SpellId,
			binding.BagSlot,
			binding.SystemId,
			binding.Image,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (b Bindings) DeletePlayer(username string) error {
	_, err := Connection.Exec("DELETE FROM bindings_players WHERE username=?", username)

	return err
}
