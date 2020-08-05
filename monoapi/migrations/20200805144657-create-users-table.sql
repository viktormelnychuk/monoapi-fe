-- +migrate Up
CREATE TABLE users
(
    id         serial,
    first_name varchar,
    last_name  varchar,
    username   varchar NOT NULL,
    password varchar NOT NULL ,
    mono_token varchar,
    mono_user_id bigint
);
