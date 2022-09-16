FROM ubuntu:20.04

ADD journal-service /

ADD docker.env .env

ADD repository/migrations /migrations

CMD ["./journal-service"]
