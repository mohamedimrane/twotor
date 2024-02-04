CREATE TABLE twoots (
    ID INTEGER PRIMARY KEY,
    contents text NOT NULL,
    created_at datetime NOT NULL,
    user_ID INTEGER NOT NULL,
    twoot_ID INTEGER,
    FOREIGN KEY (user_ID) REFERENCES users(ID),
    FOREIGN KEY (twoot_ID) REFERENCES twoots(ID)
);
