package sessionman

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/ospokemon/ospokemon"
	"golang.org/x/net/websocket"
	"taylz.io/types"
)

type Session struct {
	Username  string
	SessionId uint
	Websocket *websocket.Conn
	Expire    time.Time
	sync.Mutex
}

func (s *Session) Refresh() {
	env := ospokemon.ENV()
	s.Expire = time.Now().Add(time.Duration(types.IntString(env["sessionlife"])) * time.Second)
}

func (s *Session) Send(message string) {
	if s.Websocket != nil {
		websocket.Message.Send(s.Websocket, message)
	} else {
		ospokemon.LOG().Add("SessionId", s.SessionId).Add("Username", s.Username).Warn("session.Frame: websocket missing")
	}
}

func (s *Session) Frame() {
	if s.Websocket == nil {
		return
	}

	player := ospokemon.Players.Cache[s.Username]
	if player == nil {
		ospokemon.LOG().Add("SessionId", s.SessionId).Add("Username", s.Username).Warn("session.Frame: player missing")
		return
	}

	entity := player.GetEntity()
	universe := ospokemon.Universes.Cache[entity.UniverseId]
	if universe == nil {
		return
	}

	data := make(map[string]interface{})
	data["entityid"] = entity.Id
	data["universe"] = universe.FullFrame

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
