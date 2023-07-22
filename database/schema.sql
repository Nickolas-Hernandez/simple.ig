set client_min_messages to warning;

drop schema "public" cascade;

create schema "public";

create table "User" (
    "id"          SERIAL PRIMARY KEY,
    "username"    VARCHAR(50) NOT NULL UNIQUE,
    "email"       VARCHAR(100) NOT NULL UNIQUE,
    "password"    VARCHAR(100) NOT NULL,
    "profile_pic" VARCHAR(200),
    "bio"         TEXT,
    "created_at"  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);