package world

type View struct {
	Physics
	Contour
	Controls struct {
		Current struct {
			Name       string
			Completion float64
		}
		Abilities struct {
			Name     string
			Cooldown float64
		}
	}
	Effects struct {
		Name       string
		Completion float64
	}
}
