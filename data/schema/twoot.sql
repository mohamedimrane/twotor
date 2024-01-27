CREATE TABLE twoots (
    ID INTEGER PRIMARY KEY,
    contents text NOT NULL,
    user_ID INTEGER NOT NULL,
    FOREIGN KEY (user_ID) REFERENCES users(ID)
);
