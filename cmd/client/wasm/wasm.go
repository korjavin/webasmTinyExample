package wasm

import (
	"fmt"
	"syscall/js"
)

// drawLatency is the main function called from JavaScript to draw the latency data on the canvas.
// It receives the latency value as a float64.
func drawLatency(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		fmt.Println("No latency value provided")
		return nil
	}

	latency := args[0].Float()

	// Get the canvas and context from the DOM
	document := js.Global().Get("document")
	canvas := document.Call("getElementById", "latencyCanvas")
	context := canvas.Call("getContext", "2d")

	// Clear the canvas
	context.Call("clearRect", 0, 0, canvas.Get("width"), canvas.Get("height"))

	// Draw a rectangle representing the latency
	rectWidth := latency * 10 // Scale the latency for visualization
	rectHeight := 50
	rectX := 10
	rectY := 10

	context.Set("fillStyle", "red")
	context.Call("fillRect", rectX, rectY, rectWidth, rectHeight)

	// Add text to the canvas
	context.Set("fillStyle", "black")
	context.Set("font", "16px Arial")
	context.Call("fillText", fmt.Sprintf("Latency: %.2fms", latency*1000), 10, 100)

	return nil
}

// canvasFillRect is a wrapper function to call fillRect on canvas
func canvasFillRect(this js.Value, args []js.Value) interface{} {
	x := args[0].Int()
	y := args[1].Int()
	width := args[2].Int()
	height := args[3].Int()
	color := args[4].String()

	document := js.Global().Get("document")
	canvas := document.Call("getElementById", "latencyCanvas")
	context := canvas.Call("getContext", "2d")

	context.Set("fillStyle", color)
	context.Call("fillRect", x, y, width, height)

	return nil
}

// canvasClearRect is a wrapper function to call clearRect on canvas
func canvasClearRect(this js.Value, args []js.Value) interface{} {
	x := args[0].Int()
	y := args[1].Int()
	width := args[2].Int()
	height := args[3].Int()

	document := js.Global().Get("document")
	canvas := document.Call("getElementById", "latencyCanvas")
	context := canvas.Call("getContext", "2d")

	context.Call("clearRect", x, y, width, height)

	return nil
}

func RegisterCallbacks() {
	fmt.Println("Registering WebAssembly callbacks")
	js.Global().Set("drawLatency", js.FuncOf(drawLatency))
	js.Global().Set("canvasFillRect", js.FuncOf(canvasFillRect))
	js.Global().Set("canvasClearRect", js.FuncOf(canvasClearRect))
}
