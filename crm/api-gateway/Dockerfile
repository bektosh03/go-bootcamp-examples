FROM golang:1.19-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o api-gateway

FROM alpine

COPY --from=builder /app/api-gateway .
COPY docker.env .env

CMD ["./api-gateway"]