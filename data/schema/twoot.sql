CREATE TABLE twoots (
    ID INTEGER PRIMARY KEY,
    contents text NOT NULL,
    created_at datetime NOT NULL,
    user_ID INTEGER NOT NULL,
    FOREIGN KEY (user_ID) REFERENCES users(ID)
);
