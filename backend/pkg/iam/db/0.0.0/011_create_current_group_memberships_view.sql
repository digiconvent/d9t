create view current_group_memberships as
select 
   "user",
   "group",
   start_at
from group_has_user ghu1
where start_at <= datetime('now')
  and start_at = (
    select max(start_at)
    from group_has_user ghu2
    where ghu2."user" = ghu1."user" 
      and ghu2."group" = ghu1."group"
      and ghu2.start_at <= datetime('now')
  );

create index idx_group_has_user_lookup on group_has_user("group", start_at);
create index idx_groups_parent_lookup on groups(parent);