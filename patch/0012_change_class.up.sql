ALTER TABLE items_data RENAME TO items_data_oldcopy;
CREATE TABLE items_data (
	item INTEGER,
	key TEXT,
	VALUE TEXT);
INSERT INTO items_data (item, key, value)
	SELECT id, key, value FROM items_data_oldcopy;
DROP TABLE items_data_oldcopy;
UPDATE items SET script="set-class" WHERE id=2;
INSERT INTO items_data(item, key, value)
	VALUES (2, "class", "1");