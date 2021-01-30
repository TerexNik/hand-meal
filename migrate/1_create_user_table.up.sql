create table users
(
    id                 int  not null generated always as identity unique primary key,
    username           text not null unique,
    email              text unique,
    encrypted_password text not null,
    fio                text
)