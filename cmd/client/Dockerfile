# Stage 1: Build the Go application
FROM golang:1.24-alpine AS builder

# Set the working directory inside the container

WORKDIR /app

# Copy the source code into the container
COPY wasm .

# Build the application
ENV GOOS=js
ENV GOARCH=wasm
RUN go build -o main.wasm .

FROM nginx:latest

# Remove default nginx static content
RUN rm -rf /usr/share/nginx/html/*

# Copy the HTML, JavaScript, and WebAssembly files

COPY index.html /usr/share/nginx/html/
COPY app.js /usr/share/nginx/html/
COPY --from=builder /app/main.wasm /usr/share/nginx/html/


`

# Copy the Nginx configuration file
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Expose port 80
EXPOSE 80

# Set environment variables
ENV SERVER_HOST=server
ENV SERVER_PORT=8080

# Command to start nginx
CMD ["nginx", "-g", "daemon off;"]