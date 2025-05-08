# Start from the official Go image
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker cache
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o fireboyWatergirl ./src

# Use a smaller image for the final stage
FROM alpine:latest

# Install any runtime dependencies
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/fireboyWatergirl .

# Expose port (adjust as needed for your application)
EXPOSE 8080

# Command to run the executable
CMD ["./fireboyWatergirl"]