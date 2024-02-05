CREATE TABLE IF NOT EXISTS posts (
                                     id SERIAL PRIMARY KEY,
                                     title TEXT NOT NULL,
                                     content TEXT NOT NULL,
                                     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
INSERT INTO posts(id, title, content) VALUES (8, 'will', 'be deleted');
INSERT INTO posts(id, title, content) VALUES (9, 'will', 'be changed');
INSERT INTO posts(id, title, content) VALUES (10, 'already', 'present');
INSERT INTO posts(id, title, content) VALUES (11, 'also', 'present');