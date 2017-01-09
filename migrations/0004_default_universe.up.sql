INSERT INTO universes(id, dx, dy, private) VALUES (0, 2000, 2000, 0);

insert into entities_universes(universe, id, x, y, dx, dy) values(0, 1, 100, 100, 20, 20);
insert into entities_items(entity, universe, item, amount) values (1, 0, 1, 1);

insert into entities_universes(universe, id, x, y, dx, dy) values(0, 2, 100, 105, 20, 20);
insert into entities_items(entity, universe, item, amount) values (2, 0, 1, 1);

insert into entities_universes(universe, id, x, y, dx, dy) values(0, 3, 105, 100, 20, 20);
insert into entities_items(entity, universe, item, amount) values (3, 0, 1, 1);

insert into entities_universes(universe, id, x, y, dx, dy) values(0, 4, 100, 105, 20, 20);
insert into entities_items(entity, universe, item, amount) values (4, 0, 1, 1);
