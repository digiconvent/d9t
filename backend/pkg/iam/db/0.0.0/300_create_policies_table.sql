-- policies to the system are simply conditional permissions that need to be voted on in order to be executed
-- permissions in turn are guards for reads or writes to the system
-- e.g., let's say that in order to add new users to the d9t (iam.user.create), we want additional people to confirm
-- so nobody can go rogue, adding users without any verification. This adds a layer of trust. This trust is no distrust
-- towards the users of the system. In conventional systems, if a hacker gains access to admin rights of any system,
-- that system is completely compromised. Permissions compromised -> system compromised. Instead, we can leverage the
-- herd based trust system to validate whether or not a change should be made to the system or someone should have
-- access to sensitive data

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