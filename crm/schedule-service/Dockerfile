FROM ubuntu:20.04

ADD schedule-service /

ADD docker.env .env

ADD repository/migrations /migrations

CMD ["./schedule-service"]
