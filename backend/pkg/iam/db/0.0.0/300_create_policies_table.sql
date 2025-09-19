-- policies are simply conditional permissions that need to be voted on in order to be executed
-- permissions in turn are guards for reads/writes

create table policies (
   id uuid primary key not null,
   name varchar unique not null,
   description text,
   
   -- if this is null, no vote is required to execute an action
   -- if this = 0, it must pass unanimously
   -- if this is < 0, require 100-abs(x)% of all votes to be positive
   -- if this is > 0, require at least x positive votes (which means of all eligible voters, x must vote in favour)
   required_votes int check(required_votes >= -100 and required_votes <= 100)
);