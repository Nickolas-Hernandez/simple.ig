set client_min_messages to warning;

drop schema "public" cascade;

create schema "public";

create table "users" (
    "id"          SERIAL PRIMARY KEY,
    "username"    VARCHAR(50) NOT NULL UNIQUE,
    "email"       VARCHAR(100) NOT NULL UNIQUE,
    "password"    VARCHAR(100) NOT NULL,
    "created_at"  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);