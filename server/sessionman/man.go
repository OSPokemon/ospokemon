package sessionman

import (
	"net/http"
	"strconv"
	"time"

	"github.com/cznic/mathutil"
	"golang.org/x/net/websocket"
	"ospokemon.com"
	"ztaylor.me/cast"
	"ztaylor.me/env"
)

var Cache = make(map[uint]*Session)
var sessionIdGen, _ = mathutil.NewFC32(0, 999999, true)

func Get(account *ospokemon.Account) *Session {
	for _, session := range Cache {
		if session.Username == account.Username {
			return session
		}
	}
	return nil
}

func Add(account *ospokemon.Account) *Session {
	env := env.Global()
	session := &Session{
		Username:  account.Username,
		SessionId: uint(sessionIdGen.Next()),
		Expire:    time.Now().Add(time.Duration(cast.Int(env.Get("sessionlife"))) * time.Second),
	}

	Cache[session.SessionId] = session

	return session
}

func FromRequestCookie(r *http.Request) *Session {
	if sessionCookie, err := r.Cookie("SessionId"); err == nil {
		if sessionId, err := strconv.ParseInt(sessionCookie.Value, 10, 0); err == nil {
			return Cache[uint(sessionId)]
		}
	}

	return nil
}

func Remove(account *ospokemon.Account) {
	if session := Get(account); session != nil {
		delete(Cache, session.SessionId)
	}
}

func (session *Session) Receive() (*WebsocketMessage, error) {
	message := WebsocketMessage{}
	err := websocket.JSON.Receive(session.Websocket, &message)
	return &message, err
}

func (s *Session) WriteSessionId(w http.ResponseWriter) {
	w.Header().Set("Set-Cookie", "SessionId="+strconv.Itoa(int(s.SessionId))+"; Path=/;")
}
