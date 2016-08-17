package util

type Eventer map[string]map[int]func(...interface{})

func (e Eventer) Event(event string, eventer func(...interface{})) int {
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
	for _, eventer := range e[event] {
		if eventer != nil {
			eventer(data)
		}
	}
}

var Event = make(Eventer)
