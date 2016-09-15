CREATE TABLE actions_players (
	username TEXT,
	spellid INTEGER,
	timer INTEGER,
	FOREIGN KEY(username) REFERENCES players(username) ON DELETE CASCADE
);