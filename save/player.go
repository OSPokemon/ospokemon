package save

import (
	"github.com/ospokemon/ospokemon/event"
	"time"
)

const COMP_Player = "player"
const DEFAULT_BAG_SIZE = 6

type Player struct {
	Username   string
	Level      uint
	Experience uint
	Money      uint
	Class      uint
	BagSize    uint
	Entity     *Entity
}

func init() {
	event.On(event.AccountMake, func(args ...interface{}) {
		p := MakePlayer(args[0].(*Account).Username)
		Players[p.Username] = p
	})
	event.On(event.AccountQuery, func(args ...interface{}) {
		p := Players[args[0].(*Account).Username]
		p.Query()
	})
	event.On(event.AccountInsert, func(args ...interface{}) {
		p := Players[args[0].(*Account).Username]
		p.Insert()
	})
	event.On(event.AccountDelete, func(args ...interface{}) {
		p := Players[args[0].(*Account).Username]
		p.Delete()
	})
}

func MakePlayer(username string) *Player {
	p := &Player{
		Username: username,
		BagSize:  DEFAULT_BAG_SIZE,
		Entity:   MakeEntity(),
	}

	p.Entity.AddComponent(p)

	event.Fire(event.PlayerMake, p)

	return p
}

func GetPlayer(username string) (*Player, error) {
	if p, ok := Players[username]; p != nil {
		return p, nil
	} else if ok {
		return nil, nil
	}

	p := MakePlayer(username)
	err := p.Query()

	if err != nil {
		p = nil
	}

	Players[username] = p
	return p, err
}

func (p *Player) Id() string {
	return COMP_Player
}

func (p *Player) Update(u *Universe, e *Entity, d time.Duration) {
}

func (p *Player) Snapshot() map[string]interface{} {
	return map[string]interface{}{
		"username": p.Username,
		"level":    p.Level,
	}
}

func (p *Player) SnapshotDetail() map[string]interface{} {
	data := p.Snapshot()
	data["experience"] = p.Experience
	data["money"] = p.Money

	return data
}

func (p *Player) Query() error {
	row := Connection.QueryRow(
		"SELECT level, experience, money, class, bagsize FROM players WHERE username=?",
		p.Username,
	)

	if err := row.Scan(&p.Level, &p.Experience, &p.Money, &p.Class, &p.BagSize); err != nil {
		return err
	}

	event.Fire(event.PlayerQuery, p)

	return nil
}

func (p *Player) Insert() error {
	_, err := Connection.Exec(
		"INSERT INTO players (username, level, experience, money, class, bagsize) values (?, ?, ?, ?, ?, ?)",
		p.Username,
		p.Level,
		p.Experience,
		p.Money,
		p.Class,
		p.BagSize,
	)

	if err == nil {
		event.Fire(event.PlayerInsert, p)
	}

	return err
}

func (p *Player) Delete() error {
	_, err := Connection.Exec("DELETE FROM players WHERE username=?", p.Username)

	if err == nil {
		event.Fire(event.PlayerDelete, p)
	}

	return err
}

var Players = make(map[string]*Player)
