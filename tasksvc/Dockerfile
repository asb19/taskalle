# tasksvc/Dockerfile
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install migrate CLI
RUN apk add --no-cache curl bash git
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
RUN mv migrate /usr/local/bin/

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o tasksvc cmd/server/main.go

# Final image
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /usr/local/bin/migrate /usr/local/bin/migrate
COPY --from=builder /app/tasksvc /app/tasksvc
COPY --from=builder /app/migrations ./migrations

ENV DATABASE_URL=postgres://myuser:mypassword@postgres-task:5432/taskdb?sslmode=disable
CMD ["sh", "-c", "migrate -path ./migrations -database $DATABASE_URL up && ./tasksvc"]
