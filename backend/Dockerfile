FROM golang:1.26-alpine AS builder
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/api ./cmd/main.go

FROM alpine:3.20
WORKDIR /app

# Alpine 3.20 provides postgresql15-client (no postgresql17-client); pg_isready works against PG 17.
RUN apk add --no-cache ca-certificates curl postgresql15-client

ARG MIGRATE_VERSION=v4.18.3
RUN curl -fsSL "https://github.com/golang-migrate/migrate/releases/download/${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz" \
  | tar -xz && mv migrate /usr/local/bin/migrate && chmod +x /usr/local/bin/migrate

COPY --from=builder /out/api /app/api
COPY db/migrations /app/db/migrations
COPY entrypoint.sh /app/entrypoint.sh

RUN chmod +x /app/entrypoint.sh

EXPOSE 8080
ENTRYPOINT ["/app/entrypoint.sh"]
