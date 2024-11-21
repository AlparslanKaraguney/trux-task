# Stage 1 - Build Stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install build tools for CGO
RUN apk update && apk add --no-cache gcc musl-dev

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application with CGO enabled
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server/main.go

# Stage 2 - Final Stage
FROM alpine:latest

WORKDIR /app

# Install necessary runtime dependencies
RUN apk add --no-cache ca-certificates

# Copy the built binary and any required files
COPY --from=builder /app/server .

# Expose the port the application listens on
EXPOSE 50051

# Command to run the executable
CMD ["./server"]
