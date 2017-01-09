package save

import (
	"time"
)

type Updater interface {
	Update(*Universe, *Entity, time.Duration)
}
