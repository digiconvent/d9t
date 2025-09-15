create view permission_group_has_permissions as 
select phpg.permission_group, php.permission from policy_has_permission_group phpg
left join policy_has_permission php on phpg.policy = php.policy;