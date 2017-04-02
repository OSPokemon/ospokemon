package persistence

import (
	"ospokemon.com"
	"ospokemon.com/log"
	"time"
)

func EntitiesSpawnersSelect(universe *ospokemon.Universe) (map[uint]*ospokemon.Spawner, error) {
	spawners := make(map[uint]*ospokemon.Spawner)

	rows, err := Connection.Query(
		"SELECT entity, speed FROM entities_spawners WHERE universe=?",
		universe.Id,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entityId uint
		var speedbuff uint64
		err = rows.Scan(&entityId, &speedbuff)
		if err != nil {
			return nil, err
		}

		spawner := ospokemon.MakeSpawner()

		if t := time.Duration(speedbuff); speedbuff > 0 {
			spawner.Speed = t * time.Millisecond
		}

		spawners[entityId] = spawner
	}

	log.Add("Universe", universe.Id).Add("Spawners", spawners).Debug("entities_spawners select")

	return spawners, nil
}
