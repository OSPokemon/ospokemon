CREATE TABLE bindings_players (
	username TEXT,
	key TEXT,
	name TEXT,
	image TEXT,
	script INTEGER,
	casttime INTEGER,
	cooldown INTEGER,
	FOREIGN KEY(username) REFERENCES players(username) ON DELETE CASCADE
);