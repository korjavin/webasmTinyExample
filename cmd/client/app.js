const canvas = document.getElementById('latencyCanvas');
const ctx = canvas.getContext('2d');

const wsUrl = `ws://${window.SERVER_HOST || window.location.hostname}:${window.SERVER_PORT || '8080'}/ws`;
let ws;

async function loadWasm() {
    try {
        // Create a Go instance first (ensure it's defined by wasm_exec.js)
        if (typeof Go === 'undefined') {
            console.error('Go is not defined. Make sure wasm_exec.js is loaded first.');
            return;
        }
        
        const go = new Go();
        
        // Fetch and instantiate the WebAssembly module
        console.log('Fetching main.wasm...');
        const response = await fetch('main.wasm');
        const buffer = await response.arrayBuffer();
        console.log('Instantiating WebAssembly module...');
        const result = await WebAssembly.instantiate(buffer, go.importObject);
        console.log('Running WebAssembly module...');
        
        // Run the Go WASM instance
        go.run(result.instance);
        
        // Initialize the WebSocket connection
        ws = new WebSocket(wsUrl);

        ws.onopen = () => {
            console.log('WebSocket connected');
        };

        ws.onmessage = (event) => {
            const latency = parseFloat(event.data);
            // The drawLatency function is now available in the global scope
            if (typeof drawLatency === 'function') {
                drawLatency(latency);
            } else {
                console.error('drawLatency function not found');
            }
        };

        ws.onclose = () => {
            console.log('WebSocket disconnected');
        };

        ws.onerror = (error) => {
            console.error('WebSocket error:', error);
        };

    } catch (error) {
        console.error('Error loading WebAssembly:', error);
        console.error('Error details:', error.stack);
    }
}

loadWasm();