package loader

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ospokemon/ospokemon/objects/entities"
	"github.com/ospokemon/ospokemon/registry"
	"github.com/ospokemon/ospokemon/world"
)

func init() {
	registry.Loaders["Player"] = registry.LoaderFunc(LoadPlayer)
}

func LoadPlayer(id int) {
	if registry.Players[id] != nil {
		return
	}

	player := &entities.Player{
		STATS:   make(map[string]world.Stat),
		PHYSICS: &world.Physics{world.Position{}, world.Size{64, 64}, true},
	}

	row := Connection.QueryRow("SELECT id, name, class, x, y FROM players WHERE id=?", id)
	err := row.Scan(&player.BasicTrainer.ID, &player.NAME, &player.CLASS, &player.Physics().Position.X, &player.Physics().Position.Y)
	if err != nil {
		log.Fatal(err)
	}

	// loadPlayerPhysics(player) // handled
	loadPlayerGraphics(player)
	loadPlayerStats(player)
	loadPlayerPokemon(player)
	loadPlayerAbilities(player)

	registry.Players[id] = player
}

func UnloadPlayer(entity world.Entity) {
	player, ok := entity.(*entities.Player)

	if player == nil || !ok {
		return
	}

	_, err := Connection.Exec("UPDATE players SET class=?, x=?, y=? WHERE name=?", player.Class(), player.Physics().Position.X, player.Physics().Position.Y, player.Name())
	if err != nil {
		log.Fatal(err)
	}

	// unloadPlayerPhysics(player) // don't need
	// unloadPlayerGraphics(player) // don't need
	unloadPlayerStats(player)
	unloadPlayerPokemon(player)
	unloadPlayerAbilities(player)

	delete(registry.Players, player.Id())
}

func loadPlayerGraphics(player *entities.Player) {
	LoadAnimations("trainer", player.Class())
	animations := Animations["trainer"][player.Class()]
	player.GRAPHICS = &world.Graphics{
		Portrait:   animations[world.ANIMportrait],
		Current:    animations[world.ANIMwalk_down],
		Animations: animations,
	}
}

func loadPlayerStats(player *entities.Player) {
	rows, err := Connection.Query("SELECT stat, value, maxvalue, basemaxvalue FROM players_stats WHERE player_id=?", player.Id())
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	player.STATS = make(map[string]world.Stat)

	var name string
	var stat *entities.PlayerStat
	for rows.Next() {
		stat = &entities.PlayerStat{}
		err = rows.Scan(&name, &stat.VALUE, &stat.MAXVALUE, &stat.BASEMAXVALUE)

		player.STATS[name] = stat
	}
}

func unloadPlayerStats(player *entities.Player) {
	for name, stat := range player.Stats() {
		_, err := Connection.Exec("UPDATE players_stats SET value=?, maxvalue=?, basemaxvalue=? WHERE player_id=? AND stat=?", stat.Value(), stat.MaxValue(), stat.BaseMaxValue(), player.Id(), name)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func loadPlayerAbilities(player *entities.Player) {
	player.ABILITIES = make(map[string]*world.Ability)

	for keybinding, spell := range MakeSpellsForPlayer(player.Id()) {
		player.ABILITIES[keybinding] = &world.Ability{
			Spell: spell,
		}
	}
}

func unloadPlayerAbilities(player *entities.Player) {
	// TODO
}

func loadPlayerPokemon(player *entities.Player) {
	player.SetPokemon(fetchPokemonInPlayerBox(player.Id(), 0))
}

func unloadPlayerPokemon(player *entities.Player) {

}
