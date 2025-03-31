# WebAssembly/Go Latency Visualizer

## Project Overview

This project demonstrates the use of WebAssembly with Go to visualize network latency. The application consists of two main components:

1.  **Server:** A Go server that pings a Cloudflare server (1.1.1.1), measures the latency, calculates the average latency over a 30-second window, and serves this data via a WebSocket endpoint.
2.  **Client:** A WebAssembly application that runs in the user's browser, connects to the server via a WebSocket, receives the latency data, and draws a graph on a canvas element.

## Getting Started

### Prerequisites

*   Go (latest version)
*   Docker
*   Podman (optional, for local development)

### Running the Application

1.  **Clone the repository:**

    ```bash
    git clone <repository_url>
    cd webasmTinyExample
    ```

2.  **Build the Docker images:**

    ```bash
    docker-compose build
    ```

3.  **Start the application:**

    ```bash
    docker-compose up
    ```

4.  **Open the client in your browser:**

    Navigate to `http://localhost` (or the appropriate port if configured differently).

### Configuration

The following environment variables can be configured:

*   `SERVER_PORT`: The port the Go server listens on (default: 8080).
*   `WS_URL`: The WebSocket URL the client connects to (e.g., `ws://localhost:8080/ws`).

## Architecture

The application architecture is as follows:

```mermaid
graph LR
    A[Client (Browser)] --> B(Nginx Container);
    B --> C(index.html);
    C --> D(app.js - Loader);
    D -- Load --> E(main.wasm - WebAssembly);
    E -- WebSocket --> F(Go Server Container);
    F -- Ping --> G[Cloudflare (1.1.1.1)];
    E -- Canvas API --> H(Canvas Element);
```

## Development

### Server

The server code is located in the `cmd/server` directory.

### Client

The client code is located in the `cmd/client` directory.

## Deployment

The application can be deployed using Docker Compose. A `docker-compose.ghcr.yaml` file is provided to deploy pre-built images from `ghcr.io`.

## Contributing

Contributions are welcome! Please submit a pull request with your changes.