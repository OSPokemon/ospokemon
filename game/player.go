package game

import (
	"github.com/ospokemon/ospokemon/json"
	"github.com/ospokemon/ospokemon/part"
	"github.com/ospokemon/ospokemon/space"
	"time"
)

const DEFAULT_BAG_SIZE = 10

type Player struct {
	Username   string
	Level      uint
	Experience uint
	Money      uint
	Class      uint
	BagSize    uint
	part.Parts
}

var Players = make(map[string]*Player)

func MakePlayer() *Player {
	return &Player{
		Parts: make(part.Parts),
	}
}

func BuildPlayer(username string, bagSize uint, class *Class, entity *Entity) *Player {
	player := MakePlayer()
	player.Username = username
	player.Class = class.Id
	player.BagSize = bagSize
	player.AddPart(MakeItembag(bagSize))
	player.AddPart(BuildImaging(class.Animations))
	player.AddPart(entity)

	rect := entity.Shape.(*space.Rect)
	rect.Dimension.DX = class.Dimension.DX
	rect.Dimension.DY = class.Dimension.DY
	entity.Parts = player.Parts

	return player
}

func (p *Player) Part() string {
	return part.Player
}

func (p *Player) Update(u *Universe, e *Entity, d time.Duration) {
}

func (player *Player) Json() json.Json {
	return json.Json{
		"username": player.Username,
		"level":    player.Level,
	}
}
