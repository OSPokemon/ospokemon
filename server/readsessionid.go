package server

import (
	"net/http"
	"strconv"
)

func readSessionId(r *http.Request) int {
	if sessionCookie, err := r.Cookie("SessionId"); err == nil {
		if sessionId, err := strconv.ParseInt(sessionCookie.Value, 10, 0); err == nil {
			return int(sessionId)
		}
	}

	return 0
}
