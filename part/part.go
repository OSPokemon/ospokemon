package part

type Part interface {
	Part() string
}

type Parts map[string]Part

func (p Parts) AddPart(part Part) {
	p[part.Part()] = part
}

func (p Parts) RemovePart(part Part) {
	delete(p, part.Part())
}
