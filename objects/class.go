package objects

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/api-go"
	"github.com/ospokemon/ospokemon/engine"
)

type Class struct {
	ospokemon.BasicClass
	Graphics map[engine.AnimationType]string
	Stats    map[string]float64
}

var Classes = make(map[int]*Class)

var GetClassIds func() []int
var LoadClass func(classId int) (*Class, error)
var CreateClass func(name string) (*Class, error)
var SaveClass func(class *Class) error

func GetClass(classId int) *Class {
	if Classes[classId] == nil {
		if class, err := LoadClass(classId); err == nil {
			Classes[classId] = class
		} else {
			log.WithFields(log.Fields{
				"ClassId": classId,
				"Error":   err.Error(),
			}).Info("Class lookup failed")
		}
	}

	return Classes[classId]
}

func GetAllClasses() map[int]*Class {
	for _, id := range GetClassIds() {
		GetClass(id)
	}

	return Classes
}
