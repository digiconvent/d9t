-- multiple permissions can be included into one policy, making it easier to manage and maintain
create table policy_has_permission (
   "policy" uuid not null references policies(id) on delete cascade,
   permission varchar not null references permissions(name) on delete cascade,
   primary key ("policy", permission)
);