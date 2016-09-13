package engine

import (
	"time"
)

func Run(d time.Duration) {
	for _, universe := range Multiverse {
		universe.Update(d)
	}
}
