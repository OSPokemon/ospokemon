package server

import (
	encoder "encoding/json"
	"github.com/cznic/mathutil"
	"github.com/ospokemon/ospokemon/game"
	"github.com/ospokemon/ospokemon/json"
	"github.com/ospokemon/ospokemon/option"
	"github.com/ospokemon/ospokemon/part"
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
	return part.Session
}

func (s *Session) Update(u *game.Universe, e *game.Entity, d time.Duration) {
	p := game.Players[s.Username]

	data := make(map[string]interface{})
	universeData := make(map[string]interface{})

	data["universe"] = universeData
	data["username"] = s.Username

	data["entityid"] = p.Parts[part.Entity].(*game.Entity).Id

	for entityId, entity := range u.Entities {
		if entity == nil {
			continue
		}

		key := strconv.Itoa(int(entityId))
		_, entityData := json.Entity(entity)
		universeData[key] = entityData
	}

	menus := p.Parts[part.Menus].(game.Menus)
	if menus["player"] {
		key, playerData := json.Player(p)
		data[key] = playerData
	}
	if menus["itembag"] {
		itembag := p.Parts[part.Itembag].(*game.Itembag)
		key, itembagData := json.Itembag(itembag)
		data[key] = itembagData
	}
	if menus["actions"] {
		actions := p.Parts[part.Actions].(game.Actions)
		key, actionsData := json.Actions(actions)
		data[key] = actionsData
	}
	if menus["settings"] {
		data["settings"] = true
	}

	bindings := p.Parts[part.Bindings].(game.Bindings)
	_, bindingsData := json.Bindings(bindings)
	data["bindings"] = bindingsData

	if dialog, _ := p.Parts[part.Dialog].(*game.Dialog); dialog != nil {
		data["dialog"] = json.Dialog(dialog)
	}

	snapshot, _ := encoder.Marshal(map[string]interface{}{
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
