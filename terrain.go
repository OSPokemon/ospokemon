package ospokemon

const PARTterrain = "terrain"

type Terrain struct {
	Id        uint
	Collision bool
	Image     string
}

var terrains = make(map[uint]*Terrain)

type TerrainLink uint

func MakeTerrain(id uint) *Terrain {
	terrain := &Terrain{
		Id: id,
	}

	return terrain
}

func (t *Terrain) Part() string {
	return PARTterrain
}

func (parts Parts) GetTerrain() *Terrain {
	terrain, _ := parts[PARTterrain].(*Terrain)
	return terrain
}

func GetTerrain(id uint) (*Terrain, error) {
	if terrains[id] == nil {
		if t, err := Terrains.Select(id); err == nil {
			terrains[id] = t
		} else {
			return nil, err
		}
	}

	return terrains[id], nil
}

// persistence headers
var Terrains struct {
	Select func(uint) (*Terrain, error)
}
