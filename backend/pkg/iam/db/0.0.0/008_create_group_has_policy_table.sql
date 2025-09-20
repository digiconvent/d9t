create table group_has_policy (
   "group" uuid not null,
   policy uuid not null,
   
   primary key ("group", policy),
   foreign key ("group") references groups(id),
   foreign key (policy) references policies(id)
);

create index idx_group_has_policy_lookup on group_has_policy("group");