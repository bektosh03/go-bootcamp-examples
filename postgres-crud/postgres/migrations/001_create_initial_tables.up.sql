CREATE TABLE authors (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE books (
    id UUID PRIMARY KEY,
    title VARCHAR(120) NOT NULL,
    author_id UUID NOT NULL REFERENCES authors(id)
);
