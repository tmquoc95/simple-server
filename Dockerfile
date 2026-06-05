# Stage 1: Build the Go binary
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o myserver main.go

# Stage 2: Ship the binary in a clean, microscopic container
FROM alpine:3.19
WORKDIR /
COPY --from=builder /app/myserver /myserver
EXPOSE 8080
CMD ["/myserver"]