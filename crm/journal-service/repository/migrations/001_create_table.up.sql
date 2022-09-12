CREATE TABLE journals (
    id UUID PRIMARY KEY,
    schedule_id UUID NOT NULL,
    date DATE NOT NULL
);

CREATE TABLE journal_stats (
    journal_id UUID NOT NULL REFERENCES journals(id),
    student_id UUID NOT NULL ,
    attended BOOL NOT NULL DEFAULT TRUE,
    mark INT
);

ALTER TABLE journal_stats
ADD CONSTRAINT unique_student_mark UNIQUE (journal_id, student_id);
