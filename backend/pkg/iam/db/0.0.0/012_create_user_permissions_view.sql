create view user_permissions as
select distinct
   cgm."user",
   php.permission
from current_group_memberships cgm
join group_has_policy ghp on cgm."group" = ghp."group"
join policy_has_permission php on ghp.policy = php.policy;
