create table groups (
   id uuid primary key,
   name text not null,
   type text not null check (type in ('container', 'role', 'status')),
   parent uuid,
   description text,
   
   foreign key (parent) references groups(id),
   constraint chk_not_self_parent check (id <> parent)
);