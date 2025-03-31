package main

import (
	"fmt"
	"wasm"
)

func main() {
	fmt.Println("WebAssembly module initialized")

	// Register callbacks
	wasm.RegisterCallbacks()

	// Keep the program running
	c := make(chan struct{})
	<-c
}
