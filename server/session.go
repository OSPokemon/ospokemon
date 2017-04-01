package server

import (
	"encoding/json"
	"github.com/cznic/mathutil"
	"golang.org/x/net/websocket"
	"net/http"
	"ospokemon.com/game"
	"ospokemon.com/option"
	"strconv"
	"sync"
	"time"
)

const PARTsession = "session"

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
	return PARTsession
}

func (s *Session) Update(u *game.Universe, e *game.Entity, d time.Duration) {
	player := game.Players[s.Username]
	if player == nil {
		return
	}

	data := make(map[string]interface{})
	universeData := make(map[uint]interface{})

	data["universe"] = universeData
	data["username"] = s.Username

	data["entityid"] = e.Id

	for entityId, entity := range u.Entities {
		if entity == nil {
			continue
		}

		universeData[entityId] = entity.Json()
	}

	menus := player.GetMenus()
	if menus["player"] {
		data["player"] = player.Json()
	}
	if menus["itembag"] {
		itembag := player.GetItembag()
		data["itembag"] = itembag.Json()
	}
	if menus["actions"] {
		actions := player.GetActions()
		data["actions"] = actions.Json()
	}
	if menus["settings"] {
		data["settings"] = true
	}

	data["bindings"] = player.GetBindings().Json()

	if dialog := player.GetDialog(); dialog != nil {
		data["dialog"] = dialog.Json()
	}

	if toaster := player.GetToaster(); len(*toaster) > 0 {
		data["toaster"] = toaster.Json()
		toaster.Clear()
	}

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
