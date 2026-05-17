FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 go build -o task-cli ./cmd/main.go

FROM alpine:3.19
RUN apk add --no-cache sqlite-libs
WORKDIR /app
COPY --from=builder /app/task-cli .
ENTRYPOINT ["./task-cli"]