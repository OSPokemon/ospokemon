CREATE TABLE bindings_players (
	username TEXT,
	key TEXT,
	spellid INTEGER,
	FOREIGN KEY(username) REFERENCES players(username) ON DELETE CASCADE
);