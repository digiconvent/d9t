create table policy_has_permission (
   policy uuid not null,
   permission text not null,
   
   primary key (policy, permission),
   foreign key (policy) references policies(id),
   foreign key (permission) references permissions(permission)
);