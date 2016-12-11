CREATE TABLE accounts (
	username TEXT PRIMARY KEY,
	password TEXT,
	register INTEGER
);

CREATE TABLE universes (
	id INTEGER PRIMARY KEY,
	dx NUMBER,
	dy NUMBER
);
CREATE TABLE spells (
	id INTEGER PRIMARY KEY,
	image TEXT,
	script INTEGER,
	casttime INTEGER,
	cooldown INTEGER
);
CREATE TABLE spells_data (
	id INTEGER,
	key TEXT,
	value TEXT
);
CREATE TABLE classes (
	id INTEGER,
	dx NUMBER,
	dy NUMBER
);
CREATE TABLE animations_classes (
	class INTEGER,
	key TEXT,
	value TEXT
);
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
	dy NUMBER
);
CREATE TABLE actions_players (
	username TEXT,
	spellid INTEGER,
	timer INTEGER
);
CREATE TABLE bags_players (
	username TEXT,
	itemid INTEGER,
	timer INTEGER
);
CREATE TABLE itemslots_players (
	username TEXT,
	pos INTEGER,
	item INTEGER,
	amount INTEGER
);
CREATE TABLE bindings_players (
	username TEXT,
	key TEXT,
	spellid INTEGER,
	bagslot INTEGER,
	systemid TEXT,
	image TEXT
);
CREATE TABLE items (
	id INTEGER PRIMARY KEY,
	image TEXT,
	script INTEGER,
	casttime INTEGER,
	cooldown INTEGER,
	tradable INTEGER,
	stack INTEGER,
	value INTEGER);
CREATE TABLE items_data (
	id INTEGER,
	key TEXT,
	value TEXT);
CREATE TABLE entities_universes (
	universe INTEGER,
	id INTEGER,
	image TEXT);
CREATE TABLE entities_locations (
	entity INTEGER,
	universe INTEGER,
	x NUMBER,
	y NUMBER,
	dx NUMBER,
	dy NUMBER);
CREATE TABLE entities_items (
	entity INTEGER,
	universe INTEGER,
	item INTEGER,
	amount INTEGER);