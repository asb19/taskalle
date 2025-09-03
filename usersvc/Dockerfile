# usersvc/Dockerfile
FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache curl bash git
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
RUN mv migrate /usr/local/bin/

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o usersvc cmd/server/main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /usr/local/bin/migrate /usr/local/bin/migrate
COPY --from=builder /app/usersvc /app/usersvc
COPY --from=builder /app/migrations ./migrations

ENV DATABASE_URL=postgres://myuser:mypassword@postgres-user:5432/userdb?sslmode=disable
CMD ["sh", "-c", "migrate -path ./migrations -database $DATABASE_URL up && ./usersvc"]
