CREATE TABLE journal (
    id UUID PRIMARY KEY,
    schedule_id UUID NOT NULL,
    date DATE NOT NULL
);

CREATE TABLE journal_status (
    journal_id UUID NOT NULL REFERENCES journal(id),
    student_id UUID NOT NULL ,
    attended BOOL NOT NULL DEFAULT TRUE,
    mark INT
);

ALTER TABLE journal
ADD CONSTRAINT unique_student_mark UNIQUE (journal_id, student_id);
