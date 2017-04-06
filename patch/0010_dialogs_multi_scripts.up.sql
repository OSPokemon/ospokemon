DROP TABLE dialogs;
DROP TABLE dialogs_data;
CREATE TABLE dialogs(
	entity INTEGER,
	universe INTEGER,
	id INTEGER,
	parent INTEGER,
	lead TEXT,
	`text` TEXT);
CREATE TABLE dialogs_items_tests (
	entity INTEGER,
	universe INTEGER,
	dialog INTEGER,
	item INTEGER,
	amount INTEGER);
CREATE TABLE dialogs_scripts(
	entity INTEGER,
	universe INTEGER,
	dialog INTEGER,
	script TEXT);
CREATE TABLE dialogs_scripts_data(
	entity INTEGER,
	universe INTEGER,
	dialog INTEGER,
	script TEXT,
	key TEXT,
	value TEXT);
-- repair Ojisan dialog
INSERT INTO dialogs(entity, universe, id, parent, lead, `text`)
	VALUES (27, 0, 0, 0, "Ojisan", "Ojisan: Spare change, sir?");
INSERT INTO dialogs(entity, universe, id, parent, lead, `text`)
	VALUES (27, 0, 1, 0, "Here's 1 coin", "Ojisan: Thank you!");
INSERT INTO dialogs_items_tests(entity, universe, dialog, item, amount)
	VALUES (27, 0, 1, 1, 1);
INSERT INTO dialogs_scripts(entity, universe, dialog, script)
	VALUES (27, 0, 1, "npc-ojisan-coin");
INSERT INTO dialogs(entity, universe, id, parent, lead, `text`)
	VALUES (27, 0, 2, 0, "I'm sorry...", "Ojisan: Oh...");
