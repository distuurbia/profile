-- Create new profiles table
create table profiles (
	id uuid,
	username VARCHAR,
	password VARCHAR,
	refreshToken VARCHAR,
	country VARCHAR,
	age INTEGER,
	primary key (id)
);
