CREATE extension IF NOT EXISTS CITEXT;

DROP TABLE IF EXISTS "forum_user";
DROP TABLE IF EXISTS "vote";
DROP TABLE IF EXISTS "post";
DROP TABLE IF EXISTS "thread";
DROP TABLE IF EXISTS "forum";
DROP TABLE IF EXISTS "user";

CREATE UNLOGGED TABLE "user" (
    id serial not null UNIQUE,
    nickname CITEXT collate "C" UNIQUE not null,
    fullname TEXT not null,
    about TEXT,
    email CITEXT collate "C" not null UNIQUE
);

CREATE UNLOGGED TABLE "forum" (
    id          serial not null UNIQUE,
    title       text not null,
    "user"      citext references "user"(nickname) on delete cascade not null,
    slug        citext not null unique,
    posts       int default 0,
    threads     int default 0
);

CREATE UNLOGGED TABLE "thread" (
    id          serial not null UNIQUE,
    title       text not null,
    author      citext references "user"(nickname) on delete cascade not null,
    forum       citext references "forum"(slug) on delete cascade not null,
    message     text not null,
    votes       int default 0,
    slug        citext,
    created     timestamp with time zone default now()
);

CREATE UNLOGGED TABLE "post" (
    id          serial not null UNIQUE,
    parent      int default 0,
    author      citext references "user"(nickname) on delete cascade not null,
    message     text not null,
    is_edited   bool not null default false,
    forum       citext references "forum"(slug) on delete cascade not null,
    thread      int references "thread"(id) on delete cascade not null,
    created     timestamp with time zone default now(),
    path        int[]
);

CREATE UNLOGGED TABLE "vote" (
    id          serial not null UNIQUE,
    thread      int references "thread"(id) on delete cascade not null,
    "user"      citext references "user"(nickname) on delete cascade not null,
    voice       int not null,
    UNIQUE (thread, "user")
);

CREATE UNLOGGED TABLE "forum_user"(
    id          serial not null UNIQUE,
    forum       citext references "forum"(slug) on delete cascade not null,
    "user"      citext references "user"(nickname) on delete cascade not null,
    UNIQUE (forum, "user")
);

create or replace function update_thread_votes() returns trigger as $update_thread_votes$
begin
    update thread set votes = (votes + new.voice) where id = new.thread;
    return new;
end;
$update_thread_votes$ language plpgsql;

DROP TRIGGER IF EXISTS vote_create ON "vote";
create trigger vote_create after insert on vote for each row execute procedure update_thread_votes();

create or replace function change_thread_votes() returns trigger as $change_thread_votes$
begin
    update thread set votes = (votes - old.voice + new.voice) where id = new.thread;
    return new;
end;
$change_thread_votes$ language plpgsql;

DROP TRIGGER IF EXISTS change_thread_votes ON "vote";
create trigger change_thread_votes after update on vote for each row execute procedure change_thread_votes();

create or replace function create_post_with_path() returns trigger as $create_post$
declare
    parent_path int[];
begin
    update forum set posts = posts + 1 where slug = new.forum;
    insert into forum_user (forum, "user") values (new.forum, new.author)
    on conflict do nothing;
    parent_path = (select path from post where id = new.parent limit 1);
    new.path = parent_path || NEW.id;
    return new;
end;
$create_post$ language plpgsql;

DROP TRIGGER IF EXISTS create_post_with_path ON "post";
create trigger create_post_with_path before insert on post for each row execute procedure create_post_with_path();

create or replace function increment_forum_threads() returns trigger as $increment_forum_threads$
begin
    update forum set threads = threads + 1 where forum.slug = new.forum;
    insert into forum_user (forum, "user") values (new.forum, new.author)
    on conflict do nothing;
    return new;
end;
$increment_forum_threads$ language plpgsql;

DROP TRIGGER IF EXISTS increment_forum_threads ON "vote";
create trigger increment_forum_threads after insert on thread for each row execute procedure increment_forum_threads();


drop index if exists idx_user_on_nickname;
drop index if exists idx_user_on_email;

drop index if exists idx_forum_slug;

drop index if exists idx_thread_on_created;
drop index if exists idx_thread_on_slug;
drop index if exists idx_thread_on_forum_and_created;
drop index if exists idx_thread_on_forum;
drop index if exists idx_thread_on_id_and_forum;
drop index if exists idx_thread_on_id_and_forum;

drop index if exists idx_post_on_id;
drop index if exists idx_post_on_thread;
drop index if exists idx_post_on_thread_and_path_and_id;
drop index if exists idx_post_on_thread_and_path_and_id;
drop index if exists idx_post_on_parent;
drop index if exists idx_post_on_parent_path_and_path;
drop index if exists idx_post_on_parent_path_and_path_and_id;
drop index if exists idx_post_on_parent_and_thread;

drop index if exists idx_vote_user_thread;

drop index if exists idx_forum_user_forumusers;

create index if not exists idx_user_on_nickname on "user" using hash(nickname);
create index if not exists idx_user_on_email on "user" using hash(email);

create index if not exists idx_forum_slug on forum using hash(slug);

create index if not exists idx_thread_on_created on "thread"(created);
create index if not exists idx_thread_on_slug on "thread" using hash(slug);
create index if not exists idx_thread_on_forum_and_created on "thread" using btree(forum, created);
create index if not exists idx_thread_on_forum on thread using hash(forum);
create index if not exists idx_thread_on_id_and_forum on "thread" using btree (id, forum);
create index if not exists idx_thread_on_forum_and_created on "thread" using btree(forum, created);

create unique index if not exists idx_post_on_id ON "post" using btree (id);
create index if not exists idx_post_on_thread on post using btree (thread);
create index if not exists idx_post_on_thread_and_path_and_id on "post" using btree (thread, path, id);
create index if not exists idx_post_on_parent on "post" using btree (parent, id);
create index if not exists idx_post_on_parent_path_and_path on "post" using btree((path[1]), path);
create index if not exists idx_post_on_parent_path_and_path_and_id ON post using btree ((path[1]), path, id);
create index if not exists idx_post_on_parent_and_thread ON post using btree (parent, thread);

create index if not exists idx_vote_user_thread on "vote" using btree("user", thread);

create index if not exists idx_forum_user_forumusers on "forum_user" using btree(forum, "user");


VACUUM;
VACUUM ANALYSE;
