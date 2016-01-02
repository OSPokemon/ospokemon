package server

import (
	"github.com/cznic/mathutil"
)

var Sessions = make(map[int]*Session)
var sessionIdGen, _ = mathutil.NewFC32(0, 999999, true)

type Session struct {
	SessionId int
	ClientId  int
	Username  string
}

func CreateSession(username string) *Session {
	account := Accounts[username]
	session := &Session{sessionIdGen.Next(), 0, username}
	account.SessionId = session.SessionId
	return session
}
