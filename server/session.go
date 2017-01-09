package server

import (
	"encoding/json"
	"github.com/cznic/mathutil"
	"github.com/ospokemon/ospokemon/option"
	"github.com/ospokemon/ospokemon/part"
	"github.com/ospokemon/ospokemon/run"
	"github.com/ospokemon/ospokemon/save"
	"golang.org/x/net/websocket"
	"net/http"
	"strconv"
	"sync"
	"time"
)

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

func (s *Session) Part() string {
	return part.SESSION
}

func (s *Session) Update(u *save.Universe, e *save.Entity, d time.Duration) {
	p := save.Players[s.Username]

	data := make(map[string]interface{})
	universeData := make(map[string]interface{})

	data["universe"] = universeData
	data["username"] = s.Username

	data["entityid"] = p.Parts[part.ENTITY].(*save.Entity).Id

	for entityId, entity := range u.Entities {
		if entity == nil {
			continue
		}

		key := strconv.Itoa(int(entityId))
		_, entityData := entity.Json(true)
		universeData[key] = entityData
	}

	menus := p.Parts[part.MENUS].(*run.Menus)
	if menus.Player {
		key, playerData := p.Json(true)
		data[key] = playerData
	}
	if menus.Itembag {
		key, itembagData := p.Parts[part.ITEMBAG].(*save.Itembag).Json(true)
		data[key] = itembagData
	}
	if menus.Actions {
		key, actionsData := p.Parts[part.ACTIONS].(save.Actions).Json(true)
		data[key] = actionsData
	}

	_, bindingsData := p.Parts[part.BINDINGS].(save.Bindings).Json(true)
	data["bindings"] = bindingsData

	snapshot, _ := json.Marshal(map[string]interface{}{
		"event": "Update",
		"data":  data,
	})

	s.Send(string(snapshot))
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
