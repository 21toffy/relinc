# Build stage
FROM golang:alpine3.16 AS builder

WORKDIR /app
COPY go.mod go.sum ./
# RUN go mod download
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.12.2/migrate.linux-amd64.tar.gz | tar xvz
RUN ls


# Run stage
FROM alpine:3.16

WORKDIR /app
COPY --from=builder /app/main ./
COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .

COPY db/migration ./migration

EXPOSE 8081

CMD ["/app/main"]

ENTRYPOINT [ "/app/start.sh" ]
