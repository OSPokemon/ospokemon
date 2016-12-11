package save

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"strconv"
	"time"
)

const COMP_Actions = "save.Actions"

type Actions map[uint]*time.Duration

func init() {
	event.On(event.PlayerMake, func(args ...interface{}) {
		p := args[0].(*Player)
		actions := make(Actions)
		p.Entity.AddComponent(actions)
	})

	event.On(event.PlayerQuery, func(args ...interface{}) {
		p := args[0].(*Player)
		actions := p.Entity.Component(COMP_Actions).(Actions)
		actions.QueryPlayer(p.Username)
	})

	event.On(event.PlayerInsert, func(args ...interface{}) {
		p := args[0].(*Player)
		actions := p.Entity.Component(COMP_Actions).(Actions)
		actions.InsertPlayer(p.Username)
	})

	event.On(event.PlayerDelete, func(args ...interface{}) {
		p := args[0].(*Player)
		actions := p.Entity.Component(COMP_Actions).(Actions)
		actions.DeletePlayer(p.Username)
	})

	event.On(event.BindingDown, func(args ...interface{}) {
		p := args[0].(*Player)
		binding := args[1].(*Binding)

		if binding.SpellId > 0 {
			p.Entity.Component(COMP_Actions).(Actions).Cast(binding)
		}
	})
}

func (a Actions) clear() {
	for k := range a {
		delete(a, k)
	}
}

func (a Actions) Cast(b *Binding) {
	if a[b.SpellId] != nil {
		return
	} else if spell, err := GetSpell(b.SpellId); spell != nil {
		timer := spell.CastTime + spell.Cooldown
		a[b.SpellId] = &timer
		b.Timer = &timer
	} else if err != nil {
		logrus.Error(err.Error())
	}
}

func (a Actions) Id() string {
	return COMP_Actions
}

func (a Actions) Update(u *Universe, e *Entity, d time.Duration) {
	for spellid, timer := range a {
		if timer == nil {
			continue
		}

		if *timer < d {
			a[spellid] = nil
		}
		*timer = *timer - d
	}
}

func (a Actions) Snapshot() map[string]interface{} {
	return nil
}

func (a Actions) SnapshotDetail() map[string]interface{} {
	data := make(map[string]interface{})
	for spellid, _ := range a {

		key := strconv.Itoa(int(spellid))
		if spell, _ := GetSpell(spellid); spell != nil {
			data[key] = spell.Snapshot()
		} else {
			data[key] = spellid
		}
	}

	return data
}

func (a Actions) QueryPlayer(username string) error {
	rows, err := Connection.Query(
		"SELECT spellid, timer FROM actions_players WHERE username=?",
		username,
	)

	if err != nil {
		return err
	}

	a.clear()
	for rows.Next() {
		var spellidbuff uint
		var timerbuff uint64

		err = rows.Scan(&spellidbuff, &timerbuff)
		if err != nil {
			return err
		}

		if t := time.Duration(timerbuff); timerbuff > 0 {
			a[spellidbuff] = &t
		} else {
			a[spellidbuff] = nil
		}
	}
	rows.Close()

	return nil
}

func (a Actions) InsertPlayer(username string) error {
	for spellid, timer := range a {
		timerbuff := 0
		if timer != nil {
			timerbuff = int(*timer)
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
