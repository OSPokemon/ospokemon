package ospokemon

import "time"

type Timer time.Duration

func (t *Timer) Set(d time.Duration) {
	*t = Timer(d)
}
func (t *Timer) Duration() time.Duration {
	if t == nil {
		return 0
	}
	return time.Duration(*t)
}
func (t *Timer) Fmt() interface{} {
	if t == nil {
		return nil
	}
	return int64(*t)
}
func (t Timer) Int64() int64 {
	return int64(t)
}
