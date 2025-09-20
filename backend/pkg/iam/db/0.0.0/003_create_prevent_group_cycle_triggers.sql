create trigger prevent_group_cycle_insert
before insert on groups
for each row
when new.parent is not null
begin
   select case when exists (
      select 1 from group_has_ancestors 
      where id = new.parent and ancestor = new.id
   ) then raise(abort, 'group hierarchy cycle detected') end;
end;

create trigger prevent_group_cycle_update
before update on groups
for each row
when new.parent is not null
begin
   select case when exists (
      select 1 from group_has_ancestors 
      where id = new.parent and ancestor = new.id
   ) then raise(abort, 'group hierarchy cycle detected') end;
end;