package world

import (
	"strconv"
	"sync"
	"time"
)

var mutex = sync.Mutex{}
var Entities = make(map[int]Entity)

func Lock() {
	mutex.Lock()
}

func Unlock() {
	mutex.Unlock()
}

func AddEntity(e Entity) int {
	Lock()
	defer Unlock()
	id := unsafeAddEntity(e)
	return id
}

func unsafeAddEntity(e Entity) int {
	id := reserveEntityId()
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
	id := entityIdDispatch.next
	entityIdDispatch.next++
	return id
}

func RemoveEntity(id int) {
	Lock()
	defer Unlock()
	unsafeRemoveEntity(id)
}

func unsafeRemoveEntity(id int) {
	delete(Entities, id)
}

func Update(now time.Time) map[string]*View {
	view := make(map[string]*View)

	Lock()
	defer Unlock()

	for id, _ := range Entities {
		UpdateEntity(id, now)
	}

	for id, entity := range Entities {
		eview := MakeView(id, entity, now)
		view[strconv.Itoa(id)] = eview
	}

	return view
}
