package run

// import (
// 	"math"
// 	"time"
// )

// type Stat struct {
// 	Value     float64
// 	Max       float64
// 	TempMax   float64
// 	Regen     float64
// 	TempRegen float64
// }

// func (s *Stat) Update(u *Universe, e *Entity, c *Chart, n string, d time.Duration) {
// 	if s.Value < 1 {
// 		return
// 	}

// 	s.Value += math.Min(s.TempMax-s.Value, s.TempRegen)
// }
