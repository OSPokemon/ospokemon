package save

type Species struct {
	Id          uint
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

var Specieses = make(map[uint]*Species)

func GetSpecies(id uint) (*Species, error) {
	if s, ok := Specieses[id]; s != nil {
		return s, nil
	} else if ok {
		return nil, nil
	}

	s := MakeSpecies(id)
	err := s.Query()
	if err != nil {
		s = nil
	}

	Specieses[id] = s
	return s, err
}

func (s *Species) Query() error {
	row := Connection.QueryRow(
		"SELECT genderratio, catchfactor, hatchsteps, height, width, xpfunc FROM species WHERE id=?",
		s.Id,
	)

	var genderratiobuff float64
	if err := row.Scan(&genderratiobuff, &s.CatchFactor, &s.HatchSteps, &s.Height, &s.Width, &s.XpFunc); err != nil {
		return err
	}

	if genderratiobuff >= 0 {
		s.GenderRatio = &genderratiobuff
	}

	if err := s.queryTypes(); err != nil {
		return err
	}
	if err := s.queryMateGroups(); err != nil {
		return err
	}
	if err := s.queryLevelMoves(); err != nil {
		return err
	}
	if err := s.queryHatchMoves(); err != nil {
		return err
	}
	if err := s.queryStats(); err != nil {
		return err
	}
	if err := s.queryAnimations(); err != nil {
		return err
	}

	return nil
}

func (s *Species) queryTypes() error {
	rows, err := Connection.Query(
		"SELECT type FROM species_types WHERE species=?",
		s.Id,
	)
	if err != nil {
		return err
	}

	for rows.Next() {
		var typebuff uint
		rows.Scan(&typebuff)
		s.Types = append(s.Types, typebuff)
	}
	rows.Close()

	return nil
}

func (s *Species) queryMateGroups() error {
	rows, err := Connection.Query(
		"SELECT group FROM species_mate_groups WHERE species=?",
		s.Id,
	)
	if err != nil {
		return err
	}

	for rows.Next() {
		var groupbuff uint
		rows.Scan(&groupbuff)
		s.MateGroups = append(s.MateGroups, groupbuff)
	}
	rows.Close()

	return nil
}

func (s *Species) queryLevelMoves() error {
	rows, err := Connection.Query(
		"SELECT level, spell FROM species_level_moves WHERE species=?",
		s.Id,
	)
	if err != nil {
		return err
	}

	for rows.Next() {
		var levelbuff, spellbuff uint
		rows.Scan(&levelbuff, &spellbuff)

		if s.LevelMoves[levelbuff] == nil {
			s.LevelMoves[levelbuff] = []uint{
				spellbuff,
			}
		} else {
			s.LevelMoves[levelbuff] = append(s.LevelMoves[levelbuff], spellbuff)
		}
	}
	rows.Close()

	return nil
}

func (s *Species) queryHatchMoves() error {
	rows, err := Connection.Query(
		"SELECT spell FROM species_hatch_moves WHERE species=?",
		s.Id,
	)
	if err != nil {
		return err
	}

	for rows.Next() {
		var movebuff uint
		rows.Scan(&movebuff)
		s.HatchMoves = append(s.HatchMoves, movebuff)
	}
	rows.Close()

	return nil
}

func (s *Species) queryStats() error {
	rows, err := Connection.Query(
		"SELECT stat, value FROM species_stats WHERE species=?",
		s.Id,
	)
	if err != nil {
		return err
	}

	for rows.Next() {
		var statbuff uint
		var valuebuff float64
		rows.Scan(&statbuff, &valuebuff)
		s.Stats[statbuff] = valuebuff
	}
	rows.Close()

	return nil
}

func (s *Species) queryAnimations() error {
	rows, err := Connection.Query(
		"SELECT key, value FROM animations_species WHERE species=?",
		s.Id,
	)
	if err != nil {
		return err
	}

	for rows.Next() {
		var keybuff, valuebuff string
		rows.Scan(&keybuff, &valuebuff)
		s.Animations[keybuff] = valuebuff
	}
	rows.Close()

	return nil
}
