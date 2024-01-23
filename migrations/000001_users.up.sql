CREATE TABLE users (
    ID INTEGER PRIMARY KEY ,
    username text NOT NULL,
    email text NOT NULL,
    password text NOT NULL,
    bio text,
    display_name text
);
