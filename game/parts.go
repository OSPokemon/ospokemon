package game

type Parter interface {
	Part() string
}

type Parts map[string]Parter

func (p Parts) AddPart(part Parter) {
	p[part.Part()] = part
}

func (p Parts) RemovePart(part Parter) {
	delete(p, part.Part())
}
