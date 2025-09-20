create table group_has_user (
   "user" uuid not null,
   "group" uuid not null,
   start_at datetime not null,
   
   primary key ("user", "group", start_at),
   foreign key ("user") references users(id),
   foreign key ("group") references groups(id)
);