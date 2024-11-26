# Stage 1 - Build Stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Only copy dependencies initially
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application (static binary) with no debug information
# If you want to include debug information, remove the -s -w flags
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server ./cmd/server/main.go

# Stage 2 - Final Stage
FROM scratch

WORKDIR /app

# Copy CA certificates for HTTPS support
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the built binary
COPY --from=builder /app/server .

# Expose the port the application listens on
EXPOSE 50051

# Command to run the executable
CMD ["./server"]
