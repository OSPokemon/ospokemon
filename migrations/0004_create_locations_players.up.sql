CREATE TABLE locations_players (
	username TEXT,
	universe INTEGER,
	x NUMBER,
	y NUMBER,
	dx NUMBER,
	dy NUMBER,
	FOREIGN KEY(username) REFERENCES players(username) ON DELETE CASCADE,
	FOREIGN KEY(universe) REFERENCES universes(universe) ON DELETE CASCADE
);