INSERT INTO entities_universes(universe, id, x, y, dx, dy)
	VALUES(0, 27, 300, 300, 0, 0);
INSERT INTO classes_entities(entity, universe, class)
	VALUES (27, 0, 1);
INSERT INTO dialogs(entity, universe, id, parent, lead, `text`, script)
	VALUES (27, 0, 0, 0, "Ojisan", "Ojisan: Spare change, sir?",  "");
INSERT INTO dialogs(entity, universe, id, parent, lead, `text`, script)
	VALUES (27, 0, 1, 0, "Here's 1 coin", "Ojisan: Thank you!", "itemchange");
INSERT INTO dialogs_data(entity, universe, dialog, key, value)
	VALUES (27, 0, 1, "item", "1");
INSERT INTO dialogs_data(entity, universe, dialog, key, value)
	VALUES (27, 0, 1, "amount", "-1");
INSERT INTO dialogs(entity, universe, id, parent, lead, `text`, script)
	VALUES (27, 0, 2, 0, "I'm sorry...", "Ojisan: Oh...", "");