package main

import (
	"fmt"
	"webasmTinyExample/internal/wasm"
)

func main() {
	fmt.Println("Go Web Assembly")
	wasm.RegisterCallbacks()
	// Prevent the function from returning, which is required in a wasm module
	<-make(chan bool)
}
