CREATE extension IF NOT EXISTS CITEXT;

DROP TABLE IF EXISTS "user";
DROP TABLE IF EXISTS "forum";

CREATE TABLE "user" (
    id serial not null UNIQUE,
    nickname CITEXT collate "C" UNIQUE not null,
    fullname TEXT not null,
    about TEXT,
    email CITEXT collate "C" not null UNIQUE
);

CREATE TABLE "forum" (
    id          serial not null UNIQUE,
    title       text not null,
    "user"      citext references "user"(nickname) on delete cascade not null,
    slug        citext not null unique,
    posts       int default 0,
    threads     int default 0
);