package engine

type Events struct {
	Listeners map[string][]func(interface{})
}

type Eventer interface {
	GetListeners(name string) []func(interface{})
	On(name string, f func(interface{}))
	Fire(name string, data interface{})
}

func (e *Events) GetListeners(name string) []func(interface{}) {
	if e.Listeners == nil {
		e.Listeners = make(map[string][]func(interface{}))
	}
	if e.Listeners[name] == nil {
		e.Listeners[name] = make([]func(interface{}), 0)
	}
	return e.Listeners[name]
}

func (e *Events) On(name string, f func(interface{})) {
	e.Listeners[name] = append(e.GetListeners(name), f)
}

func (e *Events) Fire(name string, data interface{}) {
	for _, f := range e.GetListeners(name) {
		f(data)
	}
}
