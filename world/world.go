package world

import (
	"sync"
)

var World struct {
	sync.Mutex
	Entities map[int]Entity
}

func init() {
	World.Entities = make(map[int]Entity)
}

func AddEntity(e Entity) int {
	World.Lock()
	defer World.Unlock()
	id := reserveEntityId()
	World.Entities[id] = e
	return id
}

var entityIdDispatch struct {
	sync.Mutex
	next int
}

func reserveEntityId() int {
	entityIdDispatch.Lock()
	defer entityIdDispatch.Unlock()
	id := entityIdDispatch.next
	entityIdDispatch.next++
	return id
}

func RemoveEntity(id int) {
	World.Lock()
	defer World.Unlock()
	delete(World.Entities, id)
}
