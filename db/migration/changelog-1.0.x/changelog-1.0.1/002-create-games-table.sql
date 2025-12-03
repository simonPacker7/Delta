--liquibase formatted sql
--changeset Simon.Packer:1

create type game_status as enum ('abandoned', 'no possible words', 'timelimit')
go

create type game_type as enum ('online', 'private')
go

create table games (
    id uuid primary key not null ,
    start_time timestamp with time zone ,
    end_time timestamp with time zone ,
    player_one_id uuid references users(id) ,
    player_two_id uuid references users(id) ,
    winner_id uuid references users(id) ,
    word_count integer not null ,
    status game_status not null ,
    game_type game_type not null,
    game_data jsonb
)
go

create table game_players (
    game_id uuid references games(id) ,
    player_id uuid references users(id) ,
    primary key (game_id, player_id) 
)
go

create index idx_game_players_player_id on game_players(player_id)
go