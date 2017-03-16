INSERT INTO universes(id, dx, dy, private) VALUES (0, 2000, 2000, 0);

-- some trees
INSERT INTO entities_universes(universe, id, x, y, dx, dy) VALUES(0, 1, 0, -66, 48, 66);
INSERT INTO entities_terrains(entity, universe, terrain) VALUES (1, 0, 1);

INSERT INTO entities_universes(universe, id, x, y, dx, dy) VALUES(0, 2, 48, -66, 48, 66);
INSERT INTO entities_terrains(entity, universe, terrain) VALUES (2, 0, 1);

INSERT INTO entities_universes(universe, id, x, y, dx, dy) VALUES(0, 3, 96, -66, 48, 66);
INSERT INTO entities_terrains(entity, universe, terrain) VALUES (3, 0, 1);

INSERT INTO entities_universes(universe, id, x, y, dx, dy) VALUES(0, 4, 144, -66, 48, 66);
INSERT INTO entities_terrains(entity, universe, terrain) VALUES (4, 0, 1);

INSERT INTO entities_universes(universe, id, x, y, dx, dy) VALUES(0, 5, 192, -66, 48, 66);
INSERT INTO entities_terrains(entity, universe, terrain) VALUES (5, 0, 1);

INSERT INTO entities_universes(universe, id, x, y, dx, dy) VALUES(0, 6, 240, -66, 48, 66);
INSERT INTO entities_terrains(entity, universe, terrain) VALUES (6, 0, 1);

INSERT INTO entities_universes(universe, id, x, y, dx, dy) VALUES(0, 7, 288, -66, 48, 66);
INSERT INTO entities_terrains(entity, universe, terrain) VALUES (7, 0, 1);

INSERT INTO entities_universes(universe, id, x, y, dx, dy) VALUES(0, 8, 336, -66, 48, 66);
INSERT INTO entities_terrains(entity, universe, terrain) VALUES (8, 0, 1);

INSERT INTO entities_universes(universe, id, x, y, dx, dy) VALUES(0, 9, 384, -66, 48, 66);
INSERT INTO entities_terrains(entity, universe, terrain) VALUES (9, 0, 1);

INSERT INTO entities_universes(universe, id, x, y, dx, dy) VALUES(0, 10, 432, -66, 48, 66);
INSERT INTO entities_terrains(entity, universe, terrain) VALUES (10, 0, 1);

INSERT INTO entities_universes(universe, id, x, y, dx, dy) VALUES(0, 11, 480, -66, 48, 66);
INSERT INTO entities_terrains(entity, universe, terrain) VALUES (11, 0, 1);

INSERT INTO entities_universes(universe, id, x, y, dx, dy) VALUES(0, 12, -48, 0, 48, 66);
INSERT INTO entities_terrains(entity, universe, terrain) VALUES (12, 0, 1);

INSERT INTO entities_universes(universe, id, x, y, dx, dy) VALUES(0, 13, -48, 66, 48, 66);
INSERT INTO entities_terrains(entity, universe, terrain) VALUES (13, 0, 1);

INSERT INTO entities_universes(universe, id, x, y, dx, dy) VALUES(0, 14, -48, 132, 48, 66);
INSERT INTO entities_terrains(entity, universe, terrain) VALUES (14, 0, 1);

INSERT INTO entities_universes(universe, id, x, y, dx, dy) VALUES(0, 15, -48, 198, 48, 66);
INSERT INTO entities_terrains(entity, universe, terrain) VALUES (15, 0, 1);

INSERT INTO entities_universes(universe, id, x, y, dx, dy) VALUES(0, 16, -48, 264, 48, 66);
INSERT INTO entities_terrains(entity, universe, terrain) VALUES (16, 0, 1);

INSERT INTO entities_universes(universe, id, x, y, dx, dy) VALUES(0, 17, -48, 330, 48, 66);
INSERT INTO entities_terrains(entity, universe, terrain) VALUES (17, 0, 1);

INSERT INTO entities_universes(universe, id, x, y, dx, dy) VALUES(0, 18, -48, 396, 48, 66);
INSERT INTO entities_terrains(entity, universe, terrain) VALUES (18, 0, 1);

INSERT INTO entities_universes(universe, id, x, y, dx, dy) VALUES(0, 19, -48, 462, 48, 66);
INSERT INTO entities_terrains(entity, universe, terrain) VALUES (19, 0, 1);

INSERT INTO entities_universes(universe, id, x, y, dx, dy) VALUES(0, 20, -48, 528, 48, 66);
INSERT INTO entities_terrains(entity, universe, terrain) VALUES (20, 0, 1);

INSERT INTO entities_universes(universe, id, x, y, dx, dy) VALUES(0, 21, -48, 594, 48, 66);
INSERT INTO entities_terrains(entity, universe, terrain) VALUES (21, 0, 1);

INSERT INTO entities_universes(universe, id, x, y, dx, dy) VALUES(0, 22, -48, 660, 48, 66);
INSERT INTO entities_terrains(entity, universe, terrain) VALUES (22, 0, 1);

-- some coins
INSERT INTO entities_universes(universe, id, x, y, dx, dy) VALUES(0, 23, 100, 100, 0, 0);
INSERT INTO entities_items(entity, universe, item, amount) VALUES (23, 0, 1, 1);
INSERT INTO entities_spawners(entity, universe, speed) VALUES (23, 0, 10000);

INSERT INTO entities_universes(universe, id, x, y, dx, dy) VALUES(0, 24, 120, 105, 0, 0);
INSERT INTO entities_items(entity, universe, item, amount) VALUES (24, 0, 1, 1);
INSERT INTO entities_spawners(entity, universe, speed) VALUES (24, 0, 10000);

INSERT INTO entities_universes(universe, id, x, y, dx, dy) VALUES(0, 25, 140, 100, 0, 0);
INSERT INTO entities_items(entity, universe, item, amount) VALUES (25, 0, 1, 1);
INSERT INTO entities_spawners(entity, universe, speed) VALUES (25, 0, 10000);

INSERT INTO entities_universes(universe, id, x, y, dx, dy) VALUES(0, 26, 160, 105, 0, 0);
INSERT INTO entities_items(entity, universe, item, amount) VALUES (26, 0, 1, 1);
INSERT INTO entities_spawners(entity, universe, speed) VALUES (26, 0, 10000);
