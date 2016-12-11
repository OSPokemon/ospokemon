INSERT INTO universes(id, x, y, dx, dy) VALUES (0, -1000, -1000, 2000, 2000);
insert into entities_universes(universe, id, image) values(0, 1, "");
insert into entities_locations(entity, universe, x, y, dx, dy) values (1, 0, 100, 100, 20, 20);
insert into entities_items(entity, universe, item, amount) values (1, 0, 1, 1);

insert into entities_universes(universe, id, image) values(0, 2, "");
insert into entities_locations(entity, universe, x, y, dx, dy) values (2, 0, 100, 105, 20, 20);
insert into entities_items(entity, universe, item, amount) values (2, 0, 1, 1);

insert into entities_universes(universe, id, image) values(0, 3, "");
insert into entities_locations(entity, universe, x, y, dx, dy) values (3, 0, 105, 105, 20, 20);
insert into entities_items(entity, universe, item, amount) values (3, 0, 1, 1);

insert into entities_universes(universe, id, image) values(0, 4, "");
insert into entities_locations(entity, universe, x, y, dx, dy) values (4, 0, 105, 100, 20, 20);
insert into entities_items(entity, universe, item, amount) values (4, 0, 1, 1);