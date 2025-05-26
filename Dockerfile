FROM golang:1.23.4-alpine3.21 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /app/main ./cmd/main.go

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/main /app/main

CMD ["/app/main"]