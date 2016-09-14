package server

import (
	"github.com/Sirupsen/logrus"
	"github.com/cznic/mathutil"
	"github.com/ospokemon/ospokemon/util"
	"golang.org/x/net/websocket"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const EVNT_SessionExpire = "server.Session.Expire"

type Session struct {
	Username  string
	SessionId uint
	Websocket *websocket.Conn
	Expire    time.Time
	sync.Mutex
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

func (s *Session) Refresh() {
	s.Expire = time.Now().Add(time.Duration(util.OptInt("sessionlife")) * time.Second)
}

func (s *Session) Listen() {
	for s.Websocket != nil {
		var message WebsocketMessage
		err := websocket.JSON.Receive(s.Websocket, &message)

		if err != nil {
			if s.Websocket != nil {
				logrus.Warn("server.Session.listen: " + err.Error())
				util.Event.Fire(EVNT_WebsocketDisconnect, s)
			}

			return
		} else {
			go util.Event.Fire(EVNT_WebsocketMessage, s, message)
		}
	}
}

var Sessions = make(map[uint]*Session)

var sessionIdGen, _ = mathutil.NewFC32(0, 999999, true)
