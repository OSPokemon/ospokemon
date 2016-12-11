package save

import (
	"time"
)

type Updater func(*Universe, *Entity, time.Duration) map[string]interface{}
