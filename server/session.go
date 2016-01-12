package server

import (
	"github.com/cznic/mathutil"
)

type Session struct {
	SessionId int
	ClientId  int
	Username  string
}

var Sessions = make(map[int]*Session)
var sessionIdGen, _ = mathutil.NewFC32(0, 999999, true)

func CreateSession() *Session {
	return &Session{SessionId: sessionIdGen.Next()}
}
