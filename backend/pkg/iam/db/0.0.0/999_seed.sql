-- insert core permissions
insert into permissions (permission) values
('iam.users.create'),
('iam.users.read'),
('iam.users.update'),
('iam.users.delete'),
('iam.groups.create'),
('iam.groups.read'),
('iam.groups.update'),
('iam.groups.delete'),
('iam.policies.create'),
('iam.policies.read'),
('iam.policies.update'),
('iam.policies.delete');

insert into groups (id, name, type, parent, description) values
('00000000-0000-0000-0000-000000000001', 'root', 'container', null, 'root');

insert into groups (id, name, type, parent, description) values
('00000000-0000-0000-0000-000000000002', 'bootstrap admin', 'role', '00000000-0000-0000-0000-000000000001', 'initial system administrator');

insert into policies (id, name, description, votes_required) values
('00000000-0000-0000-0000-000000000003', 'full access', 'complete system access for bootstrap', 1);

insert into policy_has_permission (policy, permission)
select '00000000-0000-0000-0000-000000000003', permission from permissions;

insert into group_has_policy ("group", policy) values
('00000000-0000-0000-0000-000000000002', '00000000-0000-0000-0000-000000000003');