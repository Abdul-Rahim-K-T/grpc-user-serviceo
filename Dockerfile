# Stage 1: Build the Go binary
FROM golang:1.22.1 AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files to leverage cached layers
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire application code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o grpc-user-service ./cmd/server/main.go

# Stage 2: Create a minimal image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/grpc-user-service .

# Expose the gRPC port
EXPOSE 50051

# Command to run the binary
CMD ["./grpc-user-service"]

