--liquibase formatted sql
--changeset Simon.Packer:1

create type win_reasons as enum ('forfeit', 'no possible words', 'timelimit')
go

create type game_types as enum ('online', 'private')
go

create table games (
    id uuid primary key not null ,
    game_type game_types not null ,
    created_at timestamp with time zone ,
    start_time timestamp with time zone , -- when both players joined
    end_time timestamp with time zone ,
    player_one_id uuid references users(id) ,
    player_two_id uuid references users(id) ,
    winner_id uuid references users(id) ,
    loser_id uuid references users(id) ,
    win_reason win_reasons not null ,
    word_count integer not null ,
    moves jsonb -- array of { player_id: uuid, word: text, timestamp: timestamp }
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