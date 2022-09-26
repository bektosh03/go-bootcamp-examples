create table products(
    id uuid not null primary key,
    name varchar(50) not null unique,
    quantity int,
    price int,
    original_price int
);