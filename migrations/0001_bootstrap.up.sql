CREATE TABLE accounts (
	username TEXT PRIMARY KEY,
	password TEXT,
	register INTEGER);
CREATE TABLE universes (
	id INTEGER PRIMARY KEY,
	dx NUMBER,
	dy NUMBER,
	private INTEGER);
CREATE TABLE spells (
	id INTEGER PRIMARY KEY,
	script INTEGER,
	casttime INTEGER,
	cooldown INTEGER);
CREATE TABLE spells_data (
	id INTEGER,
	key TEXT,
	value TEXT);
CREATE TABLE classes (
	id INTEGER,
	dx NUMBER,
	dy NUMBER);
CREATE TABLE animations_classes (
	class INTEGER,
	key TEXT,
	value TEXT);
CREATE TABLE players (
	username TEXT,
	level INTEGER,
	experience INTEGER,
	money INTEGER,
	class INTEGER,
	bagsize INTEGER);
CREATE TABLE locations_players (
	username TEXT,
	universe INTEGER,
	x NUMBER,
	y NUMBER,
	dx NUMBER,
	dy NUMBER);
CREATE TABLE actions_players (
	username TEXT,
	spellid INTEGER,
	timer INTEGER);
CREATE TABLE actions_bindings_players (
	username TEXT,
	spellid INTEGER,
	key TEXT);
CREATE TABLE itembags_players (
	username TEXT,
	itemid INTEGER,
	timer INTEGER);
CREATE TABLE itemslots_players (
	username TEXT,
	pos INTEGER,
	item INTEGER,
	amount INTEGER);
CREATE TABLE bindings_players (
	username TEXT,
	key TEXT,
	systemid TEXT);
CREATE TABLE bindings_items_players (
	username TEXT,
	key TEXT,
	itemid INTEGER);
CREATE TABLE items (
	id INTEGER PRIMARY KEY,
	script INTEGER,
	casttime INTEGER,
	cooldown INTEGER,
	tradable INTEGER,
	stack INTEGER,
	value INTEGER);
CREATE TABLE animations_items (
	item INTEGER,
	key TEXT,
	value TEXT);
CREATE TABLE items_data (
	id INTEGER,
	key TEXT,
	value TEXT);
CREATE TABLE entities_universes (
	universe INTEGER,
	id INTEGER,
	x NUMBER,
	y NUMBER,
	dx NUMBER,
	dy NUMBER);
CREATE TABLE entities_items (
	entity INTEGER,
	universe INTEGER,
	item INTEGER,
	amount INTEGER);
CREATE TABLE species (
	id INTEGER PRIMARY KEY,
	genderratio NUMBER,
	catchfactor NUMBER,
	hatchsteps INTEGER,
	height NUMBER,
	width NUMBER,
	xpfunc INTEGER);
CREATE TABLE species_types (
  species INTEGER,
  type INTEGER);
CREATE TABLE species_mate_groups (
	species INTEGER,
	`group` INTEGER);
CREATE TABLE species_level_moves (
	species INTEGER,
	level INTEGER,
	spell INTEGER);
CREATE TABLE species_hatch_moves (
	species INTEGER,
	spell INTEGER);
CREATE TABLE animations_species (
	species INTEGER,
	key TEXT,
	value TEXT);
CREATE TABLE species_stats (
	species INTEGER,
	stat INTEGER,
	value NUMBER);