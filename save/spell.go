package save

import (
	"errors"
	"time"
)

type Spell struct {
	Id         uint
	ScriptId   uint
	CastTime   time.Duration
	Cooldown   time.Duration
	Animations map[string]string
	Data       map[string]string
}

func MakeSpell(id uint) *Spell {
	s := &Spell{
		Id:         id,
		Animations: make(map[string]string),
		Data:       make(map[string]string),
	}

	return s
}

func GetSpell(id uint) (*Spell, error) {
	if s, ok := Spells[id]; s != nil {
		return s, nil
	} else if ok {
		return nil, nil
	} else {
		s := MakeSpell(id)
		err := s.Query()

		if err != nil {
			s = nil
		}

		Spells[id] = s
		return s, err
	}
}

func (s *Spell) Snapshot() map[string]interface{} {
	return map[string]interface{}{
		"id":         s.Id,
		"casttime":   s.CastTime,
		"cooldown":   s.Cooldown,
		"animations": s.Animations,
	}
}

func (s *Spell) Query() error {
	row := Connection.QueryRow(
		"SELECT script, casttime, cooldown FROM spells WHERE id=?",
		s.Id,
	)

	var casttimebuff, cooldownbuff int64
	if err := row.Scan(&s.ScriptId, &casttimebuff, &cooldownbuff); err != nil {
		return err
	}

	if t := time.Duration(casttimebuff); casttimebuff > 0 {
		s.CastTime = t * time.Millisecond
	}
	if t := time.Duration(cooldownbuff); cooldownbuff > 0 {
		s.Cooldown = t * time.Millisecond
	}

	rows, err := Connection.Query(
		"SELECT key, value FROM animations_spells WHERE item=?",
		s.Id,
	)
	if err != nil {
		return err
	}

	for rows.Next() {
		var keybuff, valuebuff string
		err = rows.Scan(&keybuff, &valuebuff)
		if err != nil {
			return err
		}
		s.Animations[keybuff] = valuebuff
	}
	rows.Close()

	// TODO get spell data

	return nil
}

func (s *Spell) Insert() error {
	return errors.New("save.Spell.Insert")
}

func (s *Spell) Update() error {
	if err := s.Delete(); err != nil {
		return err
	} else if err := s.Insert(); err != nil {
		return err
	}

	return nil
}

func (s *Spell) Delete() error {
	return errors.New("save.Spell.Delete")
}

var Spells = make(map[uint]*Spell)
