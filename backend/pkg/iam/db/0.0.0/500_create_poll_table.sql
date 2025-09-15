create table polls (
   id uuid primary key not null,
   policy uuid not null references policies(id),
   author uuid not null references users(id),
   created_at integer default (strftime('%s', 'now')),
   data blob
);

create table poll_votes (
   poll uuid not null references polls(id),
   user uuid not null references users(id),
   vote boolean,
   primary key (poll, user)
);