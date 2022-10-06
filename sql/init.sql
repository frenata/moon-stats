create table if not exists
players (
id string not null primary key,
display_name string not null
);

create table if not exists
games (
id string not null primary key,
player0_id string,
player1_id string,
victory boolean
);

create table if not exists
game_units (
game_id string,
player_id string,
unit_number int,
unit_type int,
played boolean,
died boolean
);
