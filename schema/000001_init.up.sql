CREATE TABLE role_table
(
    role_id   smallserial not null primary key,
    role_name varchar     not null
);

INSERT INTO role_table (role_name)
VALUES ('Admin');
INSERT INTO role_table (role_name)
VALUES ('Student');

CREATE TABLE user_table
(
    username      varchar not null unique primary key,
    "name"        varchar not null,
    surname       varchar not null,
    patronymic    varchar,
    email         varchar not null unique,
    password_hash varchar not null,
    role_id       int     not null default 2,
    created_at    timestamptz      default now(),
    foreign key (role_id)
        references role_table (role_id) on delete cascade
);

CREATE TABLE student_table
(
    student_id    smallserial not null primary key,
    username      varchar     not null unique,
    "name"        varchar     not null,
    surname       varchar     not null,
    patronymic    varchar     default null,
    email         varchar     not null unique,
    password_hash varchar     not null,
    created_at    timestamptz default now(),
    foreign key (username)
        references user_table (username) on delete cascade on update cascade
);

CREATE TABLE admin_table
(
    admin_id      smallserial not null primary key,
    username      varchar     not null unique,
    "name"        varchar     not null,
    surname       varchar     not null,
    patronymic    varchar     default null,
    email         varchar     not null unique,
    password_hash varchar     not null,
    created_at    timestamptz default now(),
    foreign key (username)
        references user_table (username) on delete cascade on update cascade
);

CREATE OR REPLACE FUNCTION insert_user_by_role()
    returns trigger
    language plpgsql
as
$$
begin
    if new.role_id = 2 then
        insert into student_table (username, name, surname, patronymic, email, password_hash, created_at)
        values (new.username, new."name", new.surname, new.patronymic, new.email, new.password_hash, new.created_at);
    else
        insert into admin_table (username, name, surname, patronymic, email, password_hash, created_at)
        values (new.username, new."name", new.surname, new.patronymic, new.email, new.password_hash, new.created_at);
    end if;

    return new;
end;
$$;

create or replace trigger tr_sort_by_id after insert on user_table
for each row execute function insert_user_by_role();

CREATE TABLE file_table
(
    file_id     serial  not null primary key,
    username    varchar not null,
    file_chunks bytea   not null,
    foreign key (username)
        references user_table (username) on delete cascade
);

-- CREATE TABLE file_table
-- (
--     file_id     serial  not null primary key,
--     student_id  int not null,
--     file_chunks bytea   not null,
--     foreign key (student_id)
--         references student_table (student_id) on delete cascade
-- );

