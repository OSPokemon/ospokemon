CREATE TABLE players (
	username TEXT,
	level INTEGER,
	experience INTEGER,
	money INTEGER,
	FOREIGN KEY(username) REFERENCES accounts(username) ON DELETE CASCADE
);
