create table groups (
    id uuid not null primary key,
    name varchar(50) not null unique,
    main_teacher_id uuid not null
);

create table students (
    id uuid not null primary key,
    first_name varchar(30) not null,
    last_name varchar(30) not null,
    email varchar(50) not null unique,
    phone_number varchar(13) not null unique,
    level int not null,
    group_id uuid not null references groups(id)
);