CREATE TABLE if not exists connections (
    id serial primary key,
    created_at timestamp default current_timestamp,
    database_name varchar(255) not null,
    host varchar(255) not null,
    port int not null,
    username varchar(255) not null,
    password varchar(255) not null,
    common_name varchar(255) not null,
    ssl_mode bool default false
)