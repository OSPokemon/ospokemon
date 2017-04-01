package ospokemon

const PARTterrain = "terrain"

type Terrain struct {
	Id        uint
	Collision bool
	Image     string
}

var Terrains = make(map[uint]*Terrain)

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
