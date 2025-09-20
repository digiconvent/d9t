create trigger enforce_group_membership_rules
before insert on group_has_user
for each row
begin
   select case when exists (
      select 1 from groups 
      where id = new."group" and type = 'container'
   ) then raise(abort, 'container groups cannot have users') end;
end;