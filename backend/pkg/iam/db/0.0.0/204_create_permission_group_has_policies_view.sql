create view permission_group_has_policies as 
select pg.root as permission_group, phpg.policy from permission_group_has_permission_group_ancestors pg 
left join policy_has_permission_group phpg on phpg.permission_group = pg.root;