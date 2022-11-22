CREATE TABLE role_table (
    role_id smallserial not null primary key,
    role_name varchar not null
);

INSERT INTO role_table (role_name) VALUES ('Admin');
INSERT INTO role_table (role_name) VALUES ('Moderator');
INSERT INTO role_table (role_name) VALUES ('Student');

CREATE TABLE user_table (
    username varchar not null unique primary key,
    "name" varchar not null,
    surname varchar not null,
    patronymic varchar,
    email varchar not null unique,
    password_hash varchar not null,
    role_id int not null default 3,
    created_at timestamptz default now(),
    foreign key (role_id)
        references role_table (role_id) on delete cascade
);

CREATE TABLE file_table (
    file_id serial not null primary key,
    username varchar not null,
    file_url varchar not null
);

