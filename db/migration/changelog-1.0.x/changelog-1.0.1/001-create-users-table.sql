--liquibase formatted sql
--changeset Simon.Packer:1

create table users (
    id uuid primary key not null ,
    username varchar(128) not null ,
    email varchar(128) unique not null ,
    hashed_password varchar(128) not null ,
    created_at timestamp with time zone ,
    last_login timestamp with time zone
)
go

create unique index idx_users_email ON users(email)
go