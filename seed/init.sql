-- Create the root superuser
CREATE ROLE root WITH SUPERUSER LOGIN PASSWORD 'root_password';

-- Create the regular user and grant necessary privileges
GRANT ALL PRIVILEGES ON DATABASE auth_db TO auth_user;
