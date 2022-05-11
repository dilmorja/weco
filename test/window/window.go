//go:build js && wasm
// +build js,wasm

package main

import "github.com/dilmorja/weco/window"

func main() {
	// TEST: window.Closed
	window.Console.Log(window.Closed) // Expected: false
}