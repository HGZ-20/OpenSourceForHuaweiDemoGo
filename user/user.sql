-- 用户表
create table "users"(
    id serial primary key,
    name varchar(32) not null,
    email varchar(64),
    age tinyint not null default 0,
    birthday timestamp,
    member_number varchar(32),
    activated_at timestamp,
    created_at timestamp not null default now(),
    updated_at timestamp
);
