-- members of which permission_groups are eligible to vote
create table policy_has_permission_group (
   "policy" uuid not null references policies(id) on delete cascade,
   permission_group uuid not null references permission_groups(id) on delete cascade,
   primary key ("policy", permission_group)
);