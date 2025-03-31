const canvas = document.getElementById('latencyCanvas');
const ctx = canvas.getContext('2d');

const wsUrl = `ws://${window.SERVER_HOST}:${window.SERVER_PORT}/ws`;
let ws;

async function loadWasm() {
    try {
        const response = await fetch('main.wasm');
        const buffer = await response.arrayBuffer();
        const module = await WebAssembly.instantiate(buffer, {
            env: {
                // Import functions for interacting with the canvas (to be implemented in Go)
                canvasFillRect: (x, y, width, height, color) => {
                    ctx.fillStyle = color;
                    ctx.fillRect(x, y, width, height);
                },
                canvasClearRect: (x, y, width, height) => {
                    ctx.clearRect(x, y, width, height);
                }
            }
        });

        const wasmExports = module.instance.exports;

        // Initialize the WebSocket connection
        ws = new WebSocket(wsUrl);

        ws.onopen = () => {
            console.log('WebSocket connected');
        };

        ws.onmessage = (event) => {
            const latency = parseFloat(event.data);
            // Call the WebAssembly function to draw the latency data
            wasmExports.drawLatency(latency);
        };

        ws.onclose = () => {
            console.log('WebSocket disconnected');
        };

        ws.onerror = (error) => {
            console.error('WebSocket error:', error);
        };

    } catch (error) {
        console.error('Error loading WebAssembly:', error);
    }
}

loadWasm();