create table users (
   id uuid primary key,
   email text not null unique,
   first_name text not null,
   last_name text not null,
   password_hash text not null default '',
   telegram integer,
   enabled boolean not null default true,
   joined_at datetime not null
);

create unique index unique_telegram on users(telegram) where telegram is not null;