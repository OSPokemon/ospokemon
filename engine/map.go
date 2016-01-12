package engine

import (
	"sync"
	"time"
)

var Maps = make(map[string]*Map)
var nextEntityId = 1
var nextEntityIdMutex sync.Mutex

var LoadMap func(id string) *Map

type MapScript func(m *Map, now time.Time, t time.Duration)

type Map struct {
	sync.Mutex
	Name       string
	Entities   map[int]Entity
	LastUpdate time.Time
	MapScript
}

func (m *Map) AddEntity(entity Entity) int {
	m.Lock()
	defer m.Unlock()
	entityId := getNextEntityId()
	*entity.EntityId() = entityId
	m.Entities[entityId] = entity
	return entityId
}

func (m *Map) RemoveEntity(entityId int) {
	m.Lock()
	defer m.Unlock()
	delete(m.Entities, entityId)
}

func getNextEntityId() int {
	nextEntityIdMutex.Lock()
	defer nextEntityIdMutex.Unlock()
	entityId := nextEntityId
	nextEntityId++
	return entityId
}

func GetMap(id string) *Map {
	if Maps[id] == nil {
		Maps[id] = LoadMap(id)
	}

	return Maps[id]
}
