package game

type Species struct {
	Id          uint
	Name        string
	Types       []uint
	GenderRatio *float64
	CatchFactor float64
	MateGroups  []uint
	HatchSteps  uint
	Height      float64
	Width       float64
	XpFunc      uint
	LevelMoves  map[uint][]uint
	HatchMoves  []uint
	Stats       map[uint]float64
	Animations  map[string]string
}

var Specieses = make(map[uint]*Species)

func MakeSpecies(id uint) *Species {
	return &Species{
		Id:         id,
		Types:      make([]uint, 0),
		MateGroups: make([]uint, 0),
		LevelMoves: make(map[uint][]uint),
		HatchMoves: make([]uint, 0),
		Animations: make(map[string]string),
	}
}
