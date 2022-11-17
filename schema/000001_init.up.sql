CREATE TABLE user_table (
    username varchar not null unique primary key,
    "name" varchar not null,
    surname varchar not null,
    patronymic varchar,
    email varchar not null unique,
    password_hash varchar not null
);