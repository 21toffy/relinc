# Build stage
FROM golang:alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
# RUN go mod download
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine

WORKDIR /app
COPY --from=builder /app/main ./
COPY app.env .

EXPOSE 8081

CMD ["/app/main"]
