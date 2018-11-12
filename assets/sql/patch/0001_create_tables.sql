CREATE TABLE accounts (
	username VARCHAR(16) PRIMARY KEY,
	password TINYBLOB,
	register INT);
CREATE TABLE actions_bindings_players (
	username VARCHAR(16),
	spell INT,
	`key` VARCHAR(255));
CREATE TABLE actions_players (
	username VARCHAR(16),
	spell INT,
	timer INT);
CREATE TABLE animations_items (
	item INT,
	`key` VARCHAR(255),
	value VARCHAR(255));
CREATE TABLE animations_classes (
	class INT,
	`key` VARCHAR(255),
	value VARCHAR(255));
CREATE TABLE animations_species (
	species INT,
	`key` VARCHAR(255),
	value VARCHAR(255));
CREATE TABLE animations_spells (
	spell INT,
	`key` VARCHAR(255),
	value VARCHAR(255));
CREATE TABLE bindings_items_players (
	username VARCHAR(16),
	`key` VARCHAR(255),
	itemid INT);
CREATE TABLE bindings_menus_players (
	username VARCHAR(16),
	`key` VARCHAR(255),
	menu VARCHAR(255));
CREATE TABLE bindings_movements_players (
	username VARCHAR(16),
	`key` VARCHAR(255),
	direction VARCHAR(255));
CREATE TABLE classes (
	id INT,
	dx DOUBLE(6,3),
	dy DOUBLE(6,3));
CREATE TABLE dialogs (
	entity INT,
	universe INT,
	id INT,
	parent INT,
	lead VARCHAR(255),
	text VARCHAR(255));
CREATE TABLE dialogs_tests (
	entity INT,
	universe INT,
	dialog INT,
	test VARCHAR(255),
	data VARCHAR(255));
CREATE TABLE dialogs_scripts (
	entity INT,
	universe INT,
	dialog INT,
	script VARCHAR(255));
CREATE TABLE dialogs_scripts_data (
	entity INT,
	universe INT,
	dialog INT,
	script VARCHAR(255),
	`key` VARCHAR(255),
	value VARCHAR(255));
CREATE TABLE classes_entities (
	entity INT,
	universe INT,
	class INT);
CREATE TABLE entities_items (
	entity INT,
	universe INT,
	item INT,
	amount INT);
CREATE TABLE entities_spawners (
	entity INT,
	universe INT,
	speed INT);
CREATE TABLE entities_terrains (
	entity INT,
	universe INT,
	terrain INT);
CREATE TABLE entities_universes (
	universe INT,
	id INT,
	x DOUBLE(24,3),
	y DOUBLE(24,3),
	dx DOUBLE(6,3),
	dy DOUBLE(6,3));
CREATE TABLE itembags_players (
	username VARCHAR(16),
	itemid INT,
	timer INT);
CREATE TABLE items (
	id INT PRIMARY KEY,
	script VARCHAR(255),
	casttime INT,
	cooldown INT,
	tradable INT,
	stack INT,
	value INT,
	dx DOUBLE(6,3),
	dy DOUBLE(6,3));
CREATE TABLE itemslots_players (
	username VARCHAR(16),
	sort INT,
	item INT,
	amount INT);
CREATE TABLE items_data (
	item INT,
	`key` VARCHAR(255),
	value VARCHAR(255));
CREATE TABLE players (
	username VARCHAR(16),
	level INT,
	experience INT,
	money INT,
	class INT,
	universe INT,
	x DOUBLE(24,3),
	y DOUBLE(24,3));
CREATE TABLE players_stats (
	username VARCHAR(16),
	stat VARCHAR(255),
	value INT,
	base INT);
CREATE TABLE pokemon (
	id INT PRIMARY KEY,
	species INT,
	name VARCHAR(255),
	xp INT,
	level INT,
	gender VARCHAR(255),
	shiny INT);
CREATE TABLE terrain (
	id INT PRIMARY KEY,
	collision INT,
	image VARCHAR(255));
CREATE TABLE spells (
	id INT PRIMARY KEY,
	script VARCHAR(255),
	casttime INT,
	cooldown INT);
CREATE TABLE spells_data (
	spell INT,
	`key` VARCHAR(255),
	value VARCHAR(255));
CREATE TABLE species (
	id INT PRIMARY KEY,
	name VARCHAR(255),
	genderratio DOUBLE(3,3),
	catchfactor DOUBLE(3,3),
	hatchsteps INT,
	height DOUBLE(6,3),
	width DOUBLE(6,3),
	xpfunc INT);
CREATE TABLE species_hatch_moves (
	species INT,
	spell INT);
CREATE TABLE species_level_moves (
	species INT,
	level INT,
	spell INT);
CREATE TABLE species_mate_groups (
	species INT,
	`group` INT);
CREATE TABLE species_stats (
	species INT,
	stat INT,
	value INT);
CREATE TABLE species_types (
	species INT,
	type INT);
CREATE TABLE universes (
	id INT PRIMARY KEY,
	dx DOUBLE(24,3),
	dy DOUBLE(24,3),
	private INT);
