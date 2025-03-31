# Stage 1: Build the Go application
FROM golang:1.24-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Initialize a Go module if it doesn't exist
RUN if [ ! -f "go.mod" ]; then go mod init client; fi

# Create a main.go file that imports the wasm package with correct path
RUN echo 'package main

import "./wasm"

func main() {
	wasm.RegisterCallbacks()
	select {}
}' > main.go

# Make sure the wasm package is recognized
RUN go mod tidy

# Build the application
ENV GOOS=js
ENV GOARCH=wasm
RUN go build -o main.wasm

# Stage 2: Serve with Nginx
FROM nginx:latest

# Remove default nginx static content
RUN rm -rf /usr/share/nginx/html/*

# Copy the HTML, JavaScript, and WebAssembly files
COPY index.html /usr/share/nginx/html/
COPY app.js /usr/share/nginx/html/
COPY --from=builder /app/main.wasm /usr/share/nginx/html/

# Copy the Go WebAssembly support file
COPY --from=builder /usr/local/go/misc/wasm/wasm_exec.js /usr/share/nginx/html/

# Copy the Nginx configuration file
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Command to start nginx
CMD ["nginx", "-g", "daemon off;"]