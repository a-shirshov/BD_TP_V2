CREATE extension IF NOT EXISTS CITEXT;

DROP TABLE IF EXISTS "user";

CREATE TABLE "user" (
    id serial not null UNIQUE,
    nickname CITEXT collate "C" UNIQUE not null,
    fullname TEXT not null,
    about TEXT,
    email CITEXT collate "C" not null UNIQUE
);