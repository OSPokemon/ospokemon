PRAGMA foreign_keys = ON;
CREATE TABLE accounts (
	username TEXT PRIMARY KEY,
	email TEXT,
	password TEXT,
	register INTEGER
);