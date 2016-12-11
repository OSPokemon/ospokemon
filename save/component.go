package save

type Components map[string]Component

type Component interface {
	Id() string
}

func (c Components) AddComponent(comp Component) {
	c[comp.Id()] = comp
}

func (c Components) RemoveComponent(comp Component) {
	delete(c, comp.Id())
}

func (c Components) Component(key string) Component {
	return c[key]
}
