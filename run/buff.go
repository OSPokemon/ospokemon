package run

// import (
// 	"time"
// )

// type Buff struct {
// 	Name       string
// 	BonusMax   float64
// 	BonusRegen float64
// 	time.Duration
// }

// func (b *Buff) Update(u *Universe, e *Entity, c *Chart, d time.Duration) bool {
// 	stat := c.Stats[b.Name]
// 	stat.TempMax += b.BonusMax
// 	stat.TempRegen += b.BonusRegen
// 	b.Duration -= d

// 	return b.Duration > 0
// }
