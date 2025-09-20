create table policies (
   id uuid primary key,
   name text not null,
   description text,
   votes_required integer not null default 1
);