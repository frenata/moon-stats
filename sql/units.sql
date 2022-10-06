create table if not exists
units (
id int not null primary key,
name string not null,
type string
);

insert or ignore into units (id, name, type)
values (66, 'Zax Jakar', 'leader');
insert or ignore into units (id, name, type)
values (42, 'Broken Vengeance', 'ranged');
insert or ignore into units (id, name, type)
values (19, 'Aegys Defense Drone', 'support');
insert or ignore into units (id, name, type)
values (132, 'Blindsider Eztli', 'melee');
insert or ignore into units (id, name, type)
values (4, 'Crankbait', 'melee');
insert or ignore into units (id, name, type)
values (118, 'Sleeper Mine', 'summon');
