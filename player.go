package ospokemon

import (
	"ospokemon.com/json"
	"time"
)

const PARTplayer = "player"

type Player struct {
	Username   string
	Level      uint
	Experience uint
	Money      uint
	Class      uint
	Parts
}

func MakePlayer() *Player {
	return &Player{
		Parts: make(Parts),
	}
}

func BuildPlayer(username string, class *Class, entity *Entity) *Player {
	player := MakePlayer()
	player.Username = username
	player.Class = class.Id

	player.AddPart(entity)
	entity.SetClass(class)
	player.AddPart(entity.Parts[PARTimaging])

	player.AddPart(MakeItembag())
	player.AddPart(make(Bindings))
	player.AddPart(MakeMenus())
	player.AddPart(&Movement{})
	player.AddPart(MakeToaster())

	entity.Parts = player.Parts

	return player
}

func (p *Player) Part() string {
	return PARTplayer
}

func (parts Parts) GetPlayer() *Player {
	player, _ := parts[PARTplayer].(*Player)
	return player
}

func (p *Player) Update(u *Universe, e *Entity, d time.Duration) {
}

func (player *Player) Json() json.Json {
	return json.Json{
		"username": player.Username,
		"level":    player.Level,
	}
}

func GetPlayer(username string) (*Player, error) {
	if Players.Cache[username] == nil {
		if p, err := Players.Select(username); err == nil {
			Players.Cache[username] = p
		} else {
			return nil, err
		}
	}

	return Players.Cache[username], nil
}

// persistence headers
var Players = struct {
	Cache  map[string]*Player
	Select func(string) (*Player, error)
	Insert func(*Player) error
	Delete func(*Player) error
}{make(map[string]*Player), nil, nil, nil}
