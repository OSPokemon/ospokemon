package save

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"time"
)

func SpellsGet(id uint) (*Spell, error) {
	var casttimebuff, cooldownbuff int64
	s := &Spell{
		Id:   id,
		Data: make(map[string]string),
	}

	row := Connection.QueryRow(
		"SELECT image, script, casttime, cooldown FROM spells WHERE id=?",
		id,
	)

	if err := row.Scan(&s.Image, &s.ScriptId, &casttimebuff, &cooldownbuff); err != nil {
		return nil, errors.New("spellsget: " + err.Error())
	}

	if t := time.Duration(casttimebuff); casttimebuff > 0 {
		s.CastTime = t * time.Millisecond
	}
	if t := time.Duration(cooldownbuff); cooldownbuff > 0 {
		s.Cooldown = t * time.Millisecond
	}

	logrus.WithFields(logrus.Fields{
		"Id": id,
	}).Debug("save.SpellsGet")

	return s, nil
}
