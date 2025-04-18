FROM golang:1.24-alpine AS builder
RUN apk add --no-cache git
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main cmd/main.go

FROM alpine:latest
RUN apk add --no-cache ca-certificates postgresql-client wget

WORKDIR /app

COPY --from=builder /build/main     /usr/local/bin/main
COPY --from=builder /build/migrations ./migrations

ENV GOOSE_VERSION=v3.24.2
RUN wget -qO /usr/local/bin/goose \
      https://github.com/pressly/goose/releases/download/${GOOSE_VERSION}/goose_linux_x86_64 \
    && chmod +x /usr/local/bin/goose

    CMD ["sh", "-c", "goose -dir ./migrations postgres \"$DSN\" up && exec /usr/local/bin/main -env=staging --log-file=./logs/app.log"]
