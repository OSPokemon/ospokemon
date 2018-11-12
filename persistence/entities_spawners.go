package persistence

import (
	"time"

	"github.com/pkg/errors"
	"ospokemon.com"
	"ztaylor.me/log"
)

func EntitiesSpawnersSelect(universe *ospokemon.Universe) (map[uint]*ospokemon.Spawner, error) {
	spawners := make(map[uint]*ospokemon.Spawner)

	rows, err := Connection.Query(
		"SELECT entity, speed FROM entities_spawners WHERE universe=?",
		universe.Id,
	)
	if err != nil {
		return nil, errors.Wrap(err, "entities_spawners.select")
	}

	spawnerslog := make([]uint, 0)

	for rows.Next() {
		var entityId uint
		var speedbuff uint64
		err = rows.Scan(&entityId, &speedbuff)
		if err != nil {
			return nil, errors.Wrap(err, "entities_spawners.scan")
		}

		spawner := ospokemon.MakeSpawner()

		if t := time.Duration(speedbuff); speedbuff > 0 {
			spawner.Speed = t * time.Millisecond
		}

		spawners[entityId] = spawner
		spawnerslog = append(spawnerslog, entityId)
	}

	log.Add("Universe", universe.Id).Add("Spawners", spawnerslog).Debug("entities_spawners select")

	return spawners, nil
}
