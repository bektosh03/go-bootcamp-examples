CREATE TABLE schedules (
    id uuid not null primary key,
    group_id uuid not null,
    subject_id uuid not null,
    teacher_id uuid not null,
    weekday int not null,
    lesson_number int not null
);