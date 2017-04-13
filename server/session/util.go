package session

import (
	"github.com/cznic/mathutil"
	"golang.org/x/net/websocket"
	"net/http"
	"ospokemon.com"
	"ospokemon.com/log"
	"ospokemon.com/option"
	"strconv"
	"time"
)

var sessionIdGen, _ = mathutil.NewFC32(0, 999999, true)

func Add(account *ospokemon.Account) *Session {
	session := &Session{
		Username:  account.Username,
		SessionId: uint(sessionIdGen.Next()),
		Expire:    time.Now().Add(time.Duration(option.Int("sessionlife")) * time.Second),
	}

	account.AddPart(session)
	Sessions[session.SessionId] = session

	log.Add("Username", account.Username).Add("SessionId", session.SessionId).Info("session: add")
	return session
}

func Get(account *ospokemon.Account) *Session {
	session, _ := account.Parts[PARTsession].(*Session)
	return session
}

func Find(r *http.Request) *Session {
	if sessionCookie, err := r.Cookie("SessionId"); err == nil {
		if sessionId, err := strconv.ParseInt(sessionCookie.Value, 10, 0); err == nil {
			return Sessions[uint(sessionId)]
		}
	}

	return nil
}

func Remove(account *ospokemon.Account) {
	if session := Get(account); session != nil {
		log.Add("Username", account.Username).Add("SessionId", session.SessionId).Info("session: remove")
		delete(Sessions, session.SessionId)
	}
}

func (session *Session) Receive() (*WebsocketMessage, error) {
	message := WebsocketMessage{}
	err := websocket.JSON.Receive(session.Websocket, &message)
	return &message, err
}
