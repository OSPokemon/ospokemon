package save

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/part"
	"strconv"
	"time"
)

type Actions map[uint]*Action

func init() {
	event.On(event.PlayerMake, func(args ...interface{}) {
		p := args[0].(*Player)
		actions := make(Actions)
		p.AddPart(actions)
	})

	event.On(event.PlayerQuery, func(args ...interface{}) {
		p := args[0].(*Player)
		actions := p.Parts[part.ACTIONS].(Actions)
		err := actions.QueryPlayer(p.Username)

		if err != nil {
			logrus.Error(err.Error())
		}

		if len(actions) < 1 {
			for spellid, action := range MakeActionsDefault() {
				actions[spellid] = action
			}
		}
	})

	event.On(event.PlayerInsert, func(args ...interface{}) {
		p := args[0].(*Player)
		actions := p.Parts[part.ACTIONS].(Actions)
		err := actions.InsertPlayer(p.Username)

		if err != nil {
			logrus.Error(err.Error())
		}
	})

	event.On(event.PlayerDelete, func(args ...interface{}) {
		p := args[0].(*Player)
		actions := p.Parts[part.ACTIONS].(Actions)
		err := actions.DeletePlayer(p.Username)

		if err != nil {
			logrus.Error(err.Error())
		}
	})

	event.On(event.BindingDown, func(args ...interface{}) {
		// p := args[0].(*Player)
		binding := args[1].(*Binding)

		if action, ok := binding.Parts[part.ACTION].(*Action); ok {
			action.Cast(binding)
		}
	})
}

func MakeActionsDefault() Actions {
	action0 := MakeAction()
	action0.Parts[part.IMAGING].(*Imaging).Image = "/img/action/0.png"
	return Actions{
		0: action0,
	}
}

func (a Actions) clear() {
	for k := range a {
		delete(a, k)
	}
}

func (a Actions) Part() string {
	return part.ACTIONS
}

func (actions Actions) Update(u *Universe, e *Entity, d time.Duration) {
	for _, action := range actions {
		action.Update(u, e, d)
	}
}

func (a Actions) Json(expand bool) (string, map[string]interface{}) {
	data := make(map[string]interface{})

	if expand {
		for spellId, action := range a {
			key := strconv.Itoa(int(spellId))
			_, actionData := action.Json(true)
			data[key] = actionData
		}
	}

	return "actions", data
}

func (actions Actions) QueryPlayer(username string) error {
	rows, err := Connection.Query(
		"SELECT spellid, timer FROM actions_players WHERE username=?",
		username,
	)

	if err != nil {
		return err
	}

	actions.clear()
	for rows.Next() {
		action := MakeAction()
		var timerbuff uint64

		err = rows.Scan(&action.Spell, &timerbuff)
		if err != nil {
			return err
		}

		if t := time.Duration(timerbuff); timerbuff > 0 {
			action.Timer = &t
		} else {
			action.Timer = nil
		}

		actions[action.Spell] = action
	}
	rows.Close()

	event.Fire(event.ActionsPlayerQuery, username, actions)

	return nil
}

func (actions Actions) InsertPlayer(username string) error {
	for spellid, action := range actions {
		timerbuff := 0
		if action.Timer != nil {
			timerbuff = int(*action.Timer)
		}

		_, err := Connection.Exec(
			"INSERT INTO actions_players (username, spellid, timer) VALUES (?, ?, ?)",
			username,
			spellid,
			timerbuff,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (a Actions) DeletePlayer(username string) error {
	_, err := Connection.Exec("DELETE FROM actions_players WHERE username=?", username)

	return err
}
