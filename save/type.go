package save

type Type struct {
	Id     uint
	Image  string
	Strong []uint
}

var Types = make(map[uint]*Type)

func MakeType(id uint) *Type {
	return &Type{
		Id:     id,
		Strong: make([]uint, 0),
	}
}

func GetType(id uint) (*Type, error) {
	if t, ok := Types[id]; t != nil {
		return t, nil
	} else if ok {
		return nil, nil
	}

	t := MakeType(id)
	if err := t.Query(); err != nil {
		t = nil
	}

	Types[id] = t
	return t, nil
}

func (t *Type) Query() error {
	row := Connection.QueryRow(
		"SELECT image FROM types WHERE id=?",
		t.Id,
	)

	if err := row.Scan(&t.Image); err != nil {
		return err
	}

	rows, err := Connection.Query(
		"SELECT type2 FROM type_advantage WHERE type1=?",
		t.Id,
	)
	if err != nil {
		return err
	}

	for rows.Next() {
		var type2buff uint

		if err := rows.Scan(&type2buff); err != nil {
			return err
		}

		t.Strong = append(t.Strong, type2buff)
	}
	rows.Close()

	return nil
}
