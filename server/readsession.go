package server

import (
	"net/http"
	"strconv"
)

func readsession(r *http.Request) *Session {
	if sessionCookie, err := r.Cookie("SessionId"); err == nil {
		if sessionId, err := strconv.ParseInt(sessionCookie.Value, 10, 0); err == nil {
			return Sessions[uint(sessionId)]
		}
	}

	return nil
}
