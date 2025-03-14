# Stage 1: Build the Go application
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o api .

# Stage 2: Create a lightweight image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/api .
CMD [ "./api" ]