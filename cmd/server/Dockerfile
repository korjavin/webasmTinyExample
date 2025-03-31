# Stage 1: Build the Go application
FROM golang:1.24-alpine AS builder

# Set the working directory inside the container

WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -mod=vendor -o server ./

# Stage 2: Create a minimal Alpine image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/server .

# Expose the port the server listens on
EXPOSE 8080

# Set environment variables
ENV SERVER_PORT=8080

# Command to run the executable
CMD ["./server"]