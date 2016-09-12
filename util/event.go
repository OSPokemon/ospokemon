package util

const EVNT_EventFire = "util/Event.Fire"

type Eventer map[string]map[int]func(...interface{})

func (e Eventer) On(event string, eventer func(...interface{})) int {
	if e[event] == nil {
		e[event] = make(map[int]func(...interface{}))
	}

	id := len(e[event])

	e[event][id] = eventer

	return id
}

func (e Eventer) Off(event string, id int) {
	e[event][id] = nil
}

func (e Eventer) Fire(event string, data ...interface{}) {
	if event != EVNT_EventFire {
		Event.Fire(EVNT_EventFire, append([]interface{}{event}, data...)...)
	}

	for _, eventer := range e[event] {
		if eventer != nil {
			eventer(data...)
		}
	}
}

var Event = make(Eventer)
