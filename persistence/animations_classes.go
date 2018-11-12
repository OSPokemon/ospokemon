package persistence

import (
	"github.com/pkg/errors"
	"ztaylor.me/log"
)

func AnimationsClassesSelect(id uint) (map[string]string, error) {
	rows, err := Connection.Query(
		"SELECT `key`, value FROM animations_classes WHERE class=?",
		id,
	)

	if err != nil {
		return nil, errors.Wrap(err, "animations_classes.select")
	}
	defer rows.Close()

	anim := make(map[string]string)

	for rows.Next() {
		var keybuff, valuebuff string
		err = rows.Scan(&keybuff, &valuebuff)
		if err != nil {
			return nil, errors.Wrap(err, "animations_classes.scan")
		}
		anim[keybuff] = valuebuff
	}

	log.Add("ClassID", id).Add("Animations", len(anim)).Debug("animations_classes select")
	return anim, nil
}
