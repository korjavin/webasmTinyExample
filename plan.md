# WebAssembly/Go Project Plan

## Project Overview

This project aims to create a simple application that demonstrates the use of WebAssembly with Go. The application consists of two parts:

1.  **Server:** A Go server that runs in a Podman container, pings a Cloudflare server, measures latency, and calculates the average latency every 30 seconds.
2.  **Client:** A WebAssembly application that runs in the user's browser, connects to the server via a WebSocket, receives the calculated latency, and draws an image on a canvas.

## 1. Project Setup and Structure:

*   **Create Directories:** Ensure the following directories exist: `cmd/client`, `cmd/server`, `internal`.
*   **`Readme.md`:** Flesh out the `Readme.md` with a project description, setup instructions, and usage examples.
*   **Dockerfiles:** Create `Dockerfile` for both the client (Nginx) and the server (Go).
*   **`docker-compose.yml`:** Define the services (client, server) and their dependencies.
*   **`docker-compose.ghcr.yaml`:** Create a version that pulls images from `ghcr.io`.

## 2. Server-Side (Go):

*   **Implement Ping Logic:** In `cmd/server/main.go`, implement the logic to ping a Cloudflare server (e.g., `1.1.1.1`).
*   **Measure Latency:** Measure the latency of each ping.
*   **Calculate Average Latency:** Calculate the average latency over a 30-second window.
*   **WebSocket Endpoint:** Create a WebSocket endpoint that serves the average latency.
*   **Environment Variables:** Read the port from an environment variable.
*   **Dockerize:** Create a `Dockerfile` to build a minimal Alpine-based Go binary.

## 3. Client-Side (WebAssembly):

*   **HTML Setup:** Create `index.html` to host the WebAssembly application. This will include a `<canvas>` element for rendering.
*   **JavaScript Loader (`app.js`):** Create a small JavaScript file (`app.js`) to:
    *   Load the WebAssembly module (`main.wasm`).
    *   Establish a WebSocket connection to the server.
    *   Handle communication between the WebAssembly module and the browser's DOM (specifically, passing data to the canvas).
*   **WebAssembly Module (`main.wasm`):**
    *   This module, compiled from Go, will:
        *   Receive latency data from the WebSocket.
        *   Perform the calculations and drawing logic.
        *   Interact with the canvas via JavaScript functions.
*   **Build WebAssembly:** Compile the client-side Go code to WebAssembly using `GOOS=js GOARCH=wasm go build -o main.wasm`.
*   **Nginx Configuration:** Configure Nginx to serve the HTML, JavaScript, and WebAssembly files.
*   **Environment Variables:** Pass the server hostname and port as environment variables to the Nginx container.

## 4. Deployment:

*   **`docker-compose.yml`:** Define the services (client, server) and their dependencies.
*   **`docker-compose.ghcr.yaml`:** Create a version that pulls images from `ghcr.io`.
*   **GitHub Actions:** Create GitHub Actions workflows to build and push Docker images to `ghcr.io`.

## 5. Documentation:

*   **`Readme.md`:** Document the project, including setup instructions, usage, and architecture.
*   **Code Comments:** Add comments to the Go and JavaScript code to explain the logic.
*   **Logging:** Implement logging in the Go server to aid in troubleshooting.

## High-Level Diagram (Mermaid):

```mermaid
graph LR
    A[Client (Browser)] --> B(Nginx Container);
    B --> C(index.html);
    C --> D(app.js - Loader);
    D -- Load --> E(main.wasm - WebAssembly);
    E -- WebSocket --> F(Go Server Container);
    F -- Ping --> G[Cloudflare (1.1.1.1)];
    E -- Canvas API --> H(Canvas Element);