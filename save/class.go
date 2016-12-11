package save

import (
	"errors"
	"github.com/ospokemon/ospokemon/space"
)

type Class struct {
	Id        uint
	Dimension space.Vector
	Animations
}

func (c *Class) clearanimations() {
	for k := range c.Animations {
		delete(c.Animations, k)
	}
}

func MakeClass(id uint) *Class {
	c := &Class{
		Id:         id,
		Dimension:  space.Vector{},
		Animations: make(map[string]string),
	}

	return c
}

func GetClass(id uint) (*Class, error) {
	if c, ok := Classes[id]; c != nil {
		return c, nil
	} else if ok {
		return nil, nil
	} else {
		c := MakeClass(id)
		err := c.Query()

		if err != nil {
			c = nil
		}

		Classes[id] = c
		return c, err
	}
}

func (c *Class) Query() error {
	row := Connection.QueryRow(
		"SELECT dx, dy FROM classes WHERE id=?",
		c.Id,
	)

	if err := row.Scan(&c.Dimension.DX, &c.Dimension.DY); err != nil {
		return err
	}

	rows, err := Connection.Query(
		"SELECT key, value FROM animations_classes WHERE class=?",
		c.Id,
	)

	if err != nil {
		return err
	}

	c.clearanimations()
	for rows.Next() {
		var keybuff, valuebuff string
		err = rows.Scan(&keybuff, &valuebuff)
		c.Animations[keybuff] = valuebuff
	}
	rows.Close()

	return nil
}

func (c *Class) Insert() error {
	return errors.New("save.Class.Insert")
}

func (c *Class) Update() error {
	if err := c.Delete(); err != nil {
		return err
	} else if err := c.Insert(); err != nil {
		return err
	}

	return nil
}

func (c *Class) Delete() error {
	return errors.New("save.Class.Delete")
}

var Classes = make(map[uint]*Class)
