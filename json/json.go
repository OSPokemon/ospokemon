package json

import (
	"strconv"
	"time"
)

type Json map[string]interface{}

type Jsoner interface {
	Json() Json
}

func FmtDuration(t *time.Duration) interface{} {
	if t == nil {
		return nil
	}
	return int64(*t)
}

func StringUint(u uint) string {
	return strconv.FormatUint(uint64(u), 10)
}

func StringInt(i int) string {
	return strconv.FormatInt(int64(i), 10)
}
