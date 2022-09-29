FROM golang:1.19-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o email-service

FROM alpine

COPY --from=builder /app/email-service .
COPY docker.env .env

CMD ["./email-service"]
