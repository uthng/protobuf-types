DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id int NOT NULL,
    age int,
    last_name text,
    first_name text,
    created_at timestamp with time zone,
);

INSERT INTO users (age, last_name, first_name, created_at) VALUES (10, "last_name1", null, null);

INSERT INTO users (age, last_name, first_name, created_at) VALUES (null, "last_name1", null, null);
