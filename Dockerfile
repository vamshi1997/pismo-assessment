# Stage 1: Build the Go application
FROM golang:latest AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first (for dependency caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o main .

# Stage 2: Run the application in a lightweight container
FROM alpine:latest

# Install necessary dependencies (if required)
RUN apk add --no-cache ca-certificates

# Set working directory in the final container
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the port your Go application runs on
EXPOSE 8080

# Run the application
CMD ["./main"]