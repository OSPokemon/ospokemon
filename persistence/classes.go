package persistence

import (
	"github.com/pkg/errors"
	"ospokemon.com"
	"ztaylor.me/log"
)

func init() {
	ospokemon.Classes.Select = ClassesSelect
}

func ClassesSelect(id uint) (*ospokemon.Class, error) {
	row := Connection.QueryRow(
		"SELECT dx, dy FROM classes WHERE id=?",
		id,
	)

	class := ospokemon.MakeClass(id)
	err := row.Scan(&class.Dimension.DX, &class.Dimension.DY)

	if err != nil {
		return nil, errors.Wrap(err, "classes.scan")
	}

	animationsClassesBuf, err := AnimationsClassesSelect(id)
	if err != nil {
		return nil, errors.Wrap(err, "classes.scan")
	}
	for k, v := range animationsClassesBuf {
		class.Animations[k] = v
	}

	log.Add("Class", class.Id).Debug("classes select")
	return class, nil
}
