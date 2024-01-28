CREATE TABLE users (
    ID integer PRIMARY KEY,
    username text NOT NULL UNIQUE,
    email text NOT NULL UNIQUE,
    password text NOT NULL,
    bio text,
    display_name text,
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL
);
