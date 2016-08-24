package server

import (
	"github.com/cznic/mathutil"
	"github.com/ospokemon/ospokemon/save"
	"github.com/ospokemon/ospokemon/util"
	"net/http"
	"strconv"
	"time"
)

const EVNT_SessionExpire = "ospokemon/server/Session.Expire"

type Session struct {
	Username  string
	SessionId uint
	Expire    time.Time
}

func NewSession(username string) *Session {
	return &Session{
		Username:  username,
		SessionId: uint(sessionIdGen.Next()),
		Expire:    time.Now().Add(util.OptDuration("sessionlife")),
	}
}

func (s *Session) WriteSessionId(w http.ResponseWriter) {
	w.Header().Set("Set-Cookie", "SessionId="+strconv.Itoa(int(s.SessionId)))
}

func Expire(args ...interface{}) {
	sessionId := args[0].(uint)
	s := args[1].(*Session)
	Sessions[sessionId] = nil
	s.SessionId = 0
	util.Event.Fire(save.EVNT_AccountLogout, s.Username)
}

var Sessions = make(map[uint]*Session)

var sessionIdGen, _ = mathutil.NewFC32(0, 999999, true)

func init() {
	util.Event.On(EVNT_SessionExpire, Expire)
}