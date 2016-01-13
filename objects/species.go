package objects

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/api-go"
	"github.com/ospokemon/ospokemon/engine"
)

type Species struct {
	ospokemon.BasicSpecies
	GRAPHICS map[engine.AnimationType]string
}

var AllSpecies = make(map[int]*Species)

var GetSpeciesIds func() []int
var LoadSpecies func(speciesId int) (*Species, error)
var CreateSpecies func(name string) (*Species, error)
var SaveSpecies func(species *Species) error

func GetSpecies(speciesId int) *Species {
	if AllSpecies[speciesId] == nil {
		if species, err := LoadSpecies(speciesId); err == nil {
			AllSpecies[speciesId] = species
		} else {
			log.WithFields(log.Fields{
				"SpeciesId": speciesId,
				"Error":     err.Error(),
			}).Info("Species lookup failed")
		}
	}

	return AllSpecies[speciesId]
}

func GetAllSpecies() map[int]*Species {
	for _, id := range GetSpeciesIds() {
		GetSpecies(id)
	}

	return AllSpecies
}
