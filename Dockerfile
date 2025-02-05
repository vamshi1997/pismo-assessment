# Stage 1: Build the Go application
FROM golang:latest AS builder

# Set the working directory inside the container
WORKDIR /app

# Set necessary environment variables for static building
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Copy go.mod and go.sum first (for dependency caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application with static linking
RUN go build -a -installsuffix cgo -o main ./cmd/main.go


# Stage 2: Run the application in a lightweight container
FROM alpine:latest

# Install necessary dependencies (if required)
RUN apk add --no-cache ca-certificates

# Set working directory in the final container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Copy the configs directory from builder stage
COPY --from=builder /app/configs ./configs

# Expose the port your Go application runs on
EXPOSE 8080

# Run the application
CMD ["./main"]