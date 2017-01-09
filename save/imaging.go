package save

import (
	"github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/event"
	"github.com/ospokemon/ospokemon/part"
	"github.com/ospokemon/ospokemon/space"
	"math"
)

type Imaging struct {
	Image      string
	Animations map[string]string
}

func MakeImaging() *Imaging {
	return &Imaging{
		Animations: make(map[string]string),
	}
}

func init() {
	event.On(event.PlayerMake, func(args ...interface{}) {
		p := args[0].(*Player)
		imaging := MakeImaging()
		p.AddPart(imaging)
	})

	event.On(event.ActionMake, func(args ...interface{}) {
		action := args[0].(*Action)
		imaging := MakeImaging()
		action.AddPart(imaging)
	})

	event.On(event.ItemslotMake, func(args ...interface{}) {
		itemslot := args[0].(*Itemslot)
		imaging := MakeImaging()
		itemslot.AddPart(imaging)
	})

	event.On(event.PlayerQuery, func(args ...interface{}) {
		p := args[0].(*Player)
		imaging := p.Parts[part.IMAGING].(*Imaging)

		if err := imaging.QueryPlayer(p); err != nil {
			logrus.Error(err.Error())
		}
	})

	event.On(event.ItembagPlayerQuery, func(args ...interface{}) {
		//username := args[0].(string)
		itembag := args[1].(*Itembag)

		for _, itemslot := range itembag.Slots {
			if itemslot == nil {
				continue
			}

			if item, err := GetItem(itemslot.Item); item != nil {
				imaging := itemslot.Parts[part.IMAGING].(*Imaging)
				imaging.readAnimations(item.Animations)
			} else if err != nil {
				logrus.Error(err.Error())
			}
		}
	})

	event.On(event.ItemslotEntityQuery, func(args ...interface{}) {
		//universeId := args[0].(uint)
		//entityId := args[1].(uint)
		itemslot := args[2].(*Itemslot)

		if item, err := GetItem(itemslot.Item); item != nil {
			imaging := itemslot.Parts[part.IMAGING].(*Imaging)
			imaging.readAnimations(item.Animations)
		} else if err != nil {
			logrus.Error(err.Error())
		}
	})

	event.On(event.MovementUpdate, func(args ...interface{}) {
		e := args[0].(*Entity)

		if v, _ := args[1].(*space.Vector); v != nil {
			if imaging, ok := e.Parts[part.IMAGING].(*Imaging); ok {
				imaging.MovementUpdate(e, v)
			}
		}
	})
}

func (i *Imaging) clear() {
	i.Image = ""
	for k := range i.Animations {
		delete(i.Animations, k)
	}
}

func (i *Imaging) readAnimations(sample map[string]string) *Imaging {
	i.Image = sample["portrait"]
	for k, v := range sample {
		i.Animations[k] = v
	}
	return i
}

func (i *Imaging) Part() string {
	return part.IMAGING
}

func (i *Imaging) Json(expand bool) (string, map[string]interface{}) {
	return "imaging", map[string]interface{}{
		"image":      i.Image,
		"animations": i.Animations,
	}
}

func (i *Imaging) QueryPlayer(p *Player) error {
	c, err := GetClass(p.Class)

	if err != nil {
		return err
	}

	i.clear()
	for key, value := range c.Animations {
		i.Animations[key] = value
	}

	i.Image = i.Animations["portrait"]

	return nil
}

func (i *Imaging) MovementUpdate(e *Entity, v *space.Vector) {
	if v == nil {
		i.Image = i.Animations["portrait"]
	} else if slope := v.AsSlope(); slope == math.Inf(-1) {
		i.Image = i.Animations["walk-up"]
	} else if slope == math.Inf(1) {
		i.Image = i.Animations["walk-down"]
	} else if v.DX > 0 {
		i.Image = i.Animations["walk-right"]
	} else {
		i.Image = i.Animations["walk-left"]
	}
}
