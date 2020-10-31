DROP TABLE IF EXISTS users;

CREATE TABLE users (
    age int,
    lastname text,
    firstname text,
    created timestamp with time zone
);

INSERT INTO users (age, lastname, firstname, created) VALUES (10, 'last1', null, null);

INSERT INTO users (age, lastname, firstname, created) VALUES (null, 'last2', null, null);
