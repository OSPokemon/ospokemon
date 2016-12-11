package event

type Eventer map[string]map[int]func(...interface{})

var eventer = make(Eventer)

func (e Eventer) On(event string, f func(...interface{})) int {
	if e[event] == nil {
		eventer[event] = make(map[int]func(...interface{}))
	}

	id := len(eventer[event])

	eventer[event][id] = f

	return id
}

func On(event string, f func(...interface{})) int {
	return eventer.On(event, f)
}

func (e Eventer) Off(event string, id int) {
	e[event][id] = nil
}

func Off(event string, id int) {
	eventer.Off(event, id)
}

func (e Eventer) Fire(event string, data ...interface{}) {
	for _, f := range eventer[event] {
		if f != nil {
			f(data...)
		}
	}
}

func Fire(event string, data ...interface{}) {
	eventer.Fire(event, data)
}
