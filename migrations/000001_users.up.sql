CREATE TABLE users (
    ID INTEGER PRIMARY KEY,
    username text NOT NULL UNIQUE,
    email text NOT NULL UNIQUE,
    password text NOT NULL,
    bio text,
    display_name text
);
