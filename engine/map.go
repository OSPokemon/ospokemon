package engine

import (
	log "github.com/Sirupsen/logrus"
	"sync"
	"time"
)

var Maps = make(map[string]*Map)
var nextEntityId = 1
var nextEntityIdMutex sync.Mutex

var LoadMap func(mapId string) (*Map, error)

type MapScript func(m *Map, now time.Time, t time.Duration)

type Map struct {
	sync.Mutex
	Name       string
	Entities   []int
	Clients    []int
	LastUpdate time.Time
	MapScript
}

func (m *Map) AddEntity(entity Entity) int {
	m.Lock()
	defer m.Unlock()

	if *entity.EntityId() < 1 {
		entityId := getNextEntityId()
		*entity.EntityId() = entityId
		Entities[entityId] = entity
	}

	m.Entities = append(m.Entities, *entity.EntityId())

	*entity.Map() = m.Name
	return *entity.EntityId()
}

func (m *Map) RemoveEntity(entityId int) {
	m.Lock()
	defer m.Unlock()
	entity := Entities[entityId]
	*entity.Map() = ""
	for i, entityId2 := range m.Entities {
		if entityId == entityId2 {
			m.Entities = append(m.Entities[:i], m.Entities[i+1:]...)
			return
		}
	}
}

func (m *Map) AddClient(clientId int) {
	m.Lock()
	defer m.Unlock()
	m.Clients = append(m.Clients, clientId)
}

func (m *Map) RemoveClient(clientId int) {
	m.Lock()
	defer m.Unlock()
	for i, clientId2 := range m.Clients {
		if clientId == clientId2 {
			m.Clients = append(m.Clients[:i], m.Clients[i+1:]...)
			return
		}
	}
}

func getNextEntityId() int {
	nextEntityIdMutex.Lock()
	defer nextEntityIdMutex.Unlock()
	entityId := nextEntityId
	nextEntityId++
	return entityId
}

func GetMap(mapId string) *Map {
	if Maps[mapId] == nil {
		if m, err := LoadMap(mapId); err == nil {
			Maps[mapId] = m
		} else {
			log.WithFields(log.Fields{
				"MapId": mapId,
			}).Info(err.Error())
		}
	}

	return Maps[mapId]
}
