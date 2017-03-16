CREATE TABLE accounts (
	username TEXT PRIMARY KEY,
	password TEXT,
	register INTEGER);
CREATE TABLE actions_bindings_players (
	username TEXT,
	spell INTEGER,
	key TEXT);
CREATE TABLE actions_players (
	username TEXT,
	spell INTEGER,
	timer INTEGER);
CREATE TABLE animations_items (
	item INTEGER,
	key TEXT,
	value TEXT);
CREATE TABLE animations_classes (
	class INTEGER,
	key TEXT,
	value TEXT);
CREATE TABLE animations_species (
	species INTEGER,
	key TEXT,
	value TEXT);
CREATE TABLE animations_spells (
	spell INTEGER,
	key TEXT,
	value TEXT);
CREATE TABLE bindings_items_players (
	username TEXT,
	key TEXT,
	itemslot INTEGER);
CREATE TABLE bindings_menus_players (
	username TEXT,
	key TEXT,
	menu TEXT);
CREATE TABLE bindings_movements_players (
	username TEXT,
	key TEXT,
	direction TEXT);
CREATE TABLE classes (
	id INTEGER,
	dx NUMBER,
	dy NUMBER);
CREATE TABLE dialogs (
	entity INTEGER,
	universe INTEGER,
	id INTEGER,
	parent INTEGER,
	lead TEXT,
	`text` TEXT,
	script TEXT);
CREATE TABLE dialogs_data (
	entity, INTEGER,
	universe INTEGER,
	dialog INTEGER,
	key TEXT,
	value TEXT);
CREATE TABLE classes_entities (
	entity INTEGER,
	universe INTEGER,
	class INTEGER);
CREATE TABLE entities_items (
	entity INTEGER,
	universe INTEGER,
	item INTEGER,
	amount INTEGER);
CREATE TABLE entities_terrains (
	entity INTEGER,
	universe INTEGER,
	terrain INTEGER);
CREATE TABLE entities_universes (
	universe INTEGER,
	id INTEGER,
	x NUMBER,
	y NUMBER,
	dx NUMBER,
	dy NUMBER);
CREATE TABLE itembags_players (
	username TEXT,
	itemid INTEGER,
	timer INTEGER);
CREATE TABLE items (
	id INTEGER PRIMARY KEY,
	script TEXT,
	casttime INTEGER,
	cooldown INTEGER,
	tradable INTEGER,
	stack INTEGER,
	value INTEGER);
CREATE TABLE itemslots_players (
	username TEXT,
	pos INTEGER,
	item INTEGER,
	amount INTEGER);
CREATE TABLE items_data (
	id INTEGER,
	key TEXT,
	value TEXT);
CREATE TABLE players (
	username TEXT,
	level INTEGER,
	experience INTEGER,
	money INTEGER,
	class INTEGER,
	bagsize INTEGER,
	universe INTEGER,
	x NUMBER,
	y NUMBER);
CREATE TABLE players_stats (
	username TEXT,
	stat TEXT,
	value NUMBER,
	base NUMBER);
CREATE TABLE pokemon (
	id INTEGER PRIMARY KEY,
	species INTEGER,
	name TEXT,
	xp INTEGER,
	level INTEGER,
	gender TEXT,
	shiny INTEGER);
CREATE TABLE terrain (
	id INTEGER PRIMARY KEY,
	collision INTEGER,
	image string);
CREATE TABLE spells (
	id INTEGER PRIMARY KEY,
	script TEXT,
	casttime INTEGER,
	cooldown INTEGER);
CREATE TABLE spells_data (
	spell INTEGER,
	key TEXT,
	value TEXT);
CREATE TABLE species (
	id INTEGER PRIMARY KEY,
	name TEXT,
	genderratio NUMBER,
	catchfactor NUMBER,
	hatchsteps INTEGER,
	height NUMBER,
	width NUMBER,
	xpfunc INTEGER);
CREATE TABLE species_hatch_moves (
	species INTEGER,
	spell INTEGER);
CREATE TABLE species_level_moves (
	species INTEGER,
	level INTEGER,
	spell INTEGER);
CREATE TABLE species_mate_groups (
	species INTEGER,
	`group` INTEGER);
CREATE TABLE species_stats (
	species INTEGER,
	stat INTEGER,
	value NUMBER);
CREATE TABLE species_types (
	species INTEGER,
	type INTEGER);
CREATE TABLE universes (
	id INTEGER PRIMARY KEY,
	dx NUMBER,
	dy NUMBER,
	private INTEGER);