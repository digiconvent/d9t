create view group_has_ancestors as
with recursive ancestors(id, ancestor) as (
   select id, parent as ancestor 
   from groups 
   where parent is not null
   
   union all
   
   select a.id, g.parent as ancestor
   from ancestors a
   join groups g on a.ancestor = g.id
   where g.parent is not null
)
select id, ancestor from ancestors;