ALTER TABLE players RENAME TO players_oldcopy;
CREATE TABLE players (
	username TEXT,
	level INTEGER,
	experience INTEGER,
	money INTEGER,
	class INTEGER,
	universe INTEGER,
	x NUMBER,
	y NUMBER);
INSERT INTO players (username, level, experience, money, class, universe, x, y)
	SELECT username, level, experience, money, class, universe, x, y FROM players_oldcopy;
DROP TABLE players_oldcopy;
ALTER TABLE itemslots_players RENAME TO itemslots_players_oldcopy;
CREATE TABLE itemslots_players (
	username TEXT,
	sort INTEGER,
	item INTEGER,
	amount INTEGER);
INSERT INTO itemslots_players (username, sort, item, amount)
	SELECT username, pos, item, amount FROM itemslots_players_oldcopy;
DROP TABLE itemslots_players_oldcopy;
ALTER TABLE bindings_items_players RENAME TO bindings_items_players_oldcopy;
CREATE TABLE bindings_items_players (
	username TEXT,
	key TEXT,
	itemid INTEGER);
INSERT INTO bindings_items_players (username, key, itemid)
	SELECT username, key, itemslot FROM bindings_items_players_oldcopy;
DROP TABLE bindings_items_players_oldcopy;
