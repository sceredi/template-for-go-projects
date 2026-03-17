# Build stage
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal
RUN go build -o main ./cmd/app/main.go

# Final stage
FROM alpine:3.21

WORKDIR /app
COPY --from=builder /app/main .

ENTRYPOINT ["./main"]
