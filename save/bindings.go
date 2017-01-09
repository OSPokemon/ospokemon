package save

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/part"
	"time"
)

type Bindings map[string]*Binding

func init() {
	event.On(event.PlayerMake, func(args ...interface{}) {
		p := args[0].(*Player)
		bindings := make(Bindings)

		p.AddPart(bindings)
	})

	event.On(event.ActionMake, func(args ...interface{}) {
		action := args[0].(*Action)
		bindings := make(Bindings)

		action.AddPart(bindings)
	})

	event.On(event.ItemslotMake, func(args ...interface{}) {
		itemslot := args[0].(*Itemslot)
		bindings := make(Bindings)

		itemslot.AddPart(bindings)
	})

	event.On(event.PlayerQuery, func(args ...interface{}) {
		p := args[0].(*Player)
		bindings := p.Parts[part.BINDINGS].(Bindings)
		err := bindings.QueryPlayer(p.Username)

		if err != nil {
			logrus.Error(err.Error())
		}

		if len(bindings) < 1 {
			for key, binding := range MakeBindingsDefault() {
				bindings[key] = binding
			}
		}
	})

	event.On(event.ActionsPlayerQuery, func(args ...interface{}) {
		username := args[0].(string)
		actions := args[1].(Actions)

		player := Players[username]
		bindings := player.Parts[part.BINDINGS].(Bindings)

		err := bindings.QueryActionsPlayer(username, actions)

		if err != nil {
			logrus.Error(err.Error())
		}
	})

	event.On(event.ItembagPlayerQuery, func(args ...interface{}) {
		username := args[0].(string)
		itembag := args[1].(*Itembag)

		player := Players[username]
		bindings := player.Parts[part.BINDINGS].(Bindings)

		err := bindings.QueryItembagPlayer(username, itembag)

		if err != nil {
			logrus.Error(err.Error())
		}
	})

	event.On(event.PlayerInsert, func(args ...interface{}) {
		p := args[0].(*Player)

		bindings := p.Parts[part.BINDINGS].(Bindings)
		err := bindings.InsertPlayer(p.Username)

		if err != nil {
			logrus.Error(err.Error())
		}
	})

	event.On(event.ActionsPlayerInsert, func(args ...interface{}) {
		username := args[0].(string)
		actions := args[1].(Actions)

		player := Players[username]
		bindings := player.Parts[part.BINDINGS].(Bindings)

		err := bindings.InsertActionsPlayer(username, actions)

		if err != nil {
			logrus.Error(err.Error())
		}
	})

	event.On(event.ItembagPlayerInsert, func(args ...interface{}) {
		username := args[0].(string)
		itembag := args[1].(*Itembag)

		p := Players[username]
		bindings := p.Parts[part.BINDINGS].(Bindings)

		err := bindings.InsertItembagPlayer(username, itembag)

		if err != nil {
			logrus.Error(err.Error())
		}
	})

	event.On(event.PlayerDelete, func(args ...interface{}) {
		p := args[0].(*Player)

		bindings := p.Parts[part.BINDINGS].(Bindings)
		err := bindings.DeletePlayer(p.Username)

		if err != nil {
			logrus.Error(err.Error())
		}
	})

	event.On(event.ActionsPlayerDelete, func(args ...interface{}) {
		username := args[0].(string)
		p := Players[username]

		bindings := p.Parts[part.BINDINGS].(Bindings)
		err := bindings.DeleteActionsPlayer(p.Username)

		if err != nil {
			logrus.Error(err.Error())
		}
	})

	event.On(event.ItembagPlayerDelete, func(args ...interface{}) {
		username := args[0].(string)
		p := Players[username]

		bindings := p.Parts[part.BINDINGS].(Bindings)
		err := bindings.DeleteItembagPlayer(p.Username)

		if err != nil {
			logrus.Error(err.Error())
		}
	})
}

func MakeBindingsDefault() Bindings {
	return Bindings{
		"w": &Binding{Key: "w", SystemId: "walk-up", Parts: make(part.Parts)},
		"d": &Binding{Key: "d", SystemId: "walk-right", Parts: make(part.Parts)},
		"a": &Binding{Key: "a", SystemId: "walk-left", Parts: make(part.Parts)},
		"s": &Binding{Key: "s", SystemId: "walk-down", Parts: make(part.Parts)},
		"c": &Binding{
			Key:      "c",
			SystemId: "menu-player",
			Parts: part.Parts{
				part.IMAGING: &Imaging{
					Image:      "/img/ui/player.png",
					Animations: make(map[string]string),
				},
			},
		},
		"b": &Binding{
			Key:      "b",
			SystemId: "menu-itembag",
			Parts: part.Parts{
				part.IMAGING: &Imaging{
					Image:      "/img/ui/bag.png",
					Animations: make(map[string]string),
				},
			},
		},
		"v": &Binding{
			Key:      "v",
			SystemId: "menu-actions",
			Parts: part.Parts{
				part.IMAGING: &Imaging{
					Image:      "/img/ui/actions.png",
					Animations: make(map[string]string),
				},
			},
		},
		"z": &Binding{Key: "z", SystemId: "menu-bindings", Parts: make(part.Parts)},
	}
}

func (b Bindings) clear() {
	for k := range b {
		delete(b, k)
	}
}

func (b Bindings) Part() string {
	return part.BINDINGS
}

func (b Bindings) Update(u *Universe, e *Entity, d time.Duration) {
	for _, binding := range b {
		binding.Update(u, e, d)
	}
}

func (b Bindings) Json(expand bool) (string, map[string]interface{}) {
	data := make(map[string]interface{})

	if expand {
		for key, binding := range b {
			_, bindingData := binding.Json(true)
			data[key] = bindingData
		}
	}

	return "bindings", data
}

func (b Bindings) QueryPlayer(username string) error {
	rows, err := Connection.Query(
		"SELECT key, systemid FROM bindings_players WHERE username=?",
		username,
	)

	if err != nil {
		return err
	}

	for rows.Next() {
		binding := MakeBinding()

		if err = rows.Scan(&binding.Key, &binding.SystemId); err != nil {
			return err
		}

		b[binding.Key] = binding
	}
	rows.Close()

	return nil
}

func (b Bindings) QueryActionsPlayer(username string, actions Actions) error {
	rows, err := Connection.Query(
		"SELECT spellid, key FROM actions_bindings_players WHERE username=?",
		username,
	)

	if err != nil {
		return err
	}

	for rows.Next() {
		binding := MakeBinding()
		var spellidbuff uint

		if err = rows.Scan(&spellidbuff, &binding.Key); err != nil {
			return err
		}

		action := actions[spellidbuff]
		binding.AddPart(action)

		b[binding.Key] = binding
		action.Parts[part.BINDINGS].(Bindings)[binding.Key] = binding
	}

	return nil
}

func (b Bindings) QueryItembagPlayer(username string, itembag *Itembag) error {
	rows, err := Connection.Query(
		"SELECT itemid, key FROM bindings_items_players WHERE username=?",
		username,
	)

	if err != nil {
		return err
	}

	for rows.Next() {
		binding := MakeBinding()
		var itemidbuff uint

		if err = rows.Scan(&itemidbuff, &binding.Key); err != nil {
			return err
		}

		for _, itemslot := range itembag.Slots {
			if itemslot == nil {
				continue
			}

			if itemslot.Item == itemidbuff {
				itemslot.Parts[part.BINDINGS].(Bindings)[binding.Key] = binding
				binding.Parts = itemslot.Parts
			}
		}

		b[binding.Key] = binding
	}

	return nil
}

func (b Bindings) InsertPlayer(username string) error {
	for key, binding := range b {
		_, err := Connection.Exec(
			"INSERT INTO bindings_players (username, key, systemid) VALUES (?, ?, ?)",
			username,
			key,
			binding.SystemId,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (b Bindings) InsertActionsPlayer(username string, action Actions) error {
	for _, binding := range b {
		if action, ok := binding.Parts[part.ACTION].(*Action); ok {
			_, err := Connection.Exec(
				"INSERT INTO actions_bindings_players (username, spellid, key) VALUES (?, ?, ?)",
				username,
				action.Spell,
				binding.Key,
			)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (b Bindings) InsertItembagPlayer(username string, itembag *Itembag) error {
	for _, binding := range b {
		if itemslot, ok := binding.Parts[part.ITEMSLOT].(*Itemslot); ok {
			_, err := Connection.Exec(
				"INSERT INTO bindings_items_players (username, itemid, key) VALUES (?, ?, ?)",
				username,
				itemslot.Item,
				binding.Key)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (b Bindings) DeletePlayer(username string) error {
	_, err := Connection.Exec("DELETE FROM bindings_players WHERE username=?", username)

	return err
}

func (b Bindings) DeleteActionsPlayer(username string) error {
	_, err := Connection.Exec("DELETE FROM actions_bindings_players WHERE username=?", username)

	return err
}

func (b Bindings) DeleteItembagPlayer(username string) error {
	_, err := Connection.Exec("DELETE FROM bindings_items_players WHERE username=?", username)

	return err
}
