drop user if exists ronly1;
drop user if exists ronly2;
create role foundations_readonly WITH PASSWORD 'ronly'; --LOGIN 
grant connect on database foundations to foundations_readonly;
grant usage on schema public to foundations_readonly;

grant select on all tables in schema public to foundations_readonly;


--create user ronly1 with password 'ronly1';
--grant foundations_readonly to ronly1;

--create user ronly2 with password 'ronly2';
--grant foundations_readonly to ronly2;
