package db

func GetClassIds() []int {
	rows, err := Connection.Query("SELECT id FROM classes")
	if err != nil {
		return nil
	}

	ids := make([]int, 0)
	for rows.Next() {
		var id int
		rows.Scan(&id)
		ids = append(ids, id)
	}

	return ids
}
