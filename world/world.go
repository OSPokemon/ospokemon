package world

import (
	log "github.com/Sirupsen/logrus"
	"sync"
)

var Entities = make(map[int]Entity)

func AddEntity(e Entity) int {
	id := reserveEntityId()
	e.SetEntityId(id)
	Entities[id] = e
	return id
}

var entityIdDispatch struct {
	sync.Mutex
	next int
}

func reserveEntityId() int {
	entityIdDispatch.Lock()
	defer entityIdDispatch.Unlock()
	entityIdDispatch.next++
	id := entityIdDispatch.next

	log.WithFields(log.Fields{
		"EntityId": id,
	}).Debug("EntityId Reserved")

	return id
}

func RemoveEntity(id int) {

	log.WithFields(log.Fields{
		"EntityId": id,
	}).Debug("EntityId Cleared")

	delete(Entities, id)
}
