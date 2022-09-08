create table journal(
    id uuid not null primary key ,
    schedule_id uuid not null ,
    student_id uuid not null ,
    attended bool not null ,
    mark int not null ,
);