package ospokemon

import (
	"ospokemon.com/json"
	"ospokemon.com/space"
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

var players = make(map[string]*Player)

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
	player.AddPart(MakeItembag())
	player.AddPart(make(Bindings))
	player.AddPart(MakeMenus())
	player.AddPart(&Movement{})
	player.AddPart(BuildImaging(class.Animations))
	player.AddPart(MakeToaster())

	rect := entity.Shape.(*space.Rect)
	rect.Dimension.DX = class.Dimension.DX
	rect.Dimension.DY = class.Dimension.DY
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
	if players[username] == nil {
		if p, err := Players.Select(username); err == nil {
			players[username] = p
		} else {
			return nil, err
		}
	}

	return players[username], nil
}

// persistence headers
var Players struct {
	Select func(string) (*Player, error)
	Insert func(*Player) error
	Delete func(*Player) error
}
