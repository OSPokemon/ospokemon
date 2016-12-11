package server

import (
	"encoding/json"
	"github.com/cznic/mathutil"
	"github.com/ospokemon/ospokemon/option"
	"github.com/ospokemon/ospokemon/run"
	"github.com/ospokemon/ospokemon/save"
	"golang.org/x/net/websocket"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const COMP_Session = "server.Session"

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
		Expire:    time.Now().Add(time.Duration(option.Int("sessionlife")) * time.Second),
	}
}

func (s *Session) Id() string {
	return COMP_Session
}

func (s *Session) Update(u *save.Universe, e *save.Entity, d time.Duration) {
	p := save.Players[s.Username]

	data := make(map[string]interface{})

	data["username"] = p.Username
	data["universe"] = u.Snapshot()
	data["bindings"] = p.Entity.Component(save.COMP_Bindings).(save.Bindings).SnapshotDetail()

	menus := p.Entity.Component(run.COMP_Menus).(*run.Menus)
	if menus.Player {
		data["player"] = p.SnapshotDetail()
	}
	if menus.Bag {
		data["bag"] = p.Entity.Component(save.COMP_Bag).(*save.Bag).SnapshotDetail()
	}
	if menus.Actions {
		data["actions"] = p.Entity.Component(save.COMP_Actions).(save.Actions).SnapshotDetail()
	}

	// data["animations"] = p.Entity.Component(COMP_Animations).(Animations).Snapshot()

	snapshot, _ := json.Marshal(map[string]interface{}{
		"event": "Update",
		"data":  data,
	})

	s.Send(string(snapshot))
}

func (s *Session) Snapshot() map[string]interface{} {
	return nil
}

func (s *Session) WriteSessionId(w http.ResponseWriter) {
	w.Header().Set("Set-Cookie", "SessionId="+strconv.Itoa(int(s.SessionId))+"; Path=/;")
}

func (s *Session) Refresh() {
	s.Expire = time.Now().Add(time.Duration(option.Int("sessionlife")) * time.Second)
}

func (s *Session) Send(message string) {
	websocket.Message.Send(s.Websocket, message)
}

var Sessions = make(map[uint]*Session)

var sessionIdGen, _ = mathutil.NewFC32(0, 999999, true)
