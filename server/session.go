package server

import (
	"github.com/cznic/mathutil"
	"github.com/ospokemon/ospokemon/util"
	"net/http"
	"strconv"
	"time"
)

const EVNT_SessionExpire = "server/Session.Expire"

type Session struct {
	Username  string
	SessionId uint
	Expire    time.Time
}

func NewSession(username string) *Session {
	return &Session{
		Username:  username,
		SessionId: uint(sessionIdGen.Next()),
		Expire:    time.Now().Add(time.Duration(util.OptInt("sessionlife")) * time.Second),
	}
}

func (s *Session) WriteSessionId(w http.ResponseWriter) {
	w.Header().Set("Set-Cookie", "SessionId="+strconv.Itoa(int(s.SessionId))+"; Path=/;")
}

var Sessions = make(map[uint]*Session)

var sessionIdGen, _ = mathutil.NewFC32(0, 999999, true)
