package world

type AnimationType uint8

const (
	ANIMwalk_down AnimationType = iota
	ANIMwalk_right
	ANIMwalk_up
	ANIMwalk_left
	ANIMcast
	ANIMstun
)

type Graphics struct {
	Current    string
	Animations map[AnimationType]string
}
