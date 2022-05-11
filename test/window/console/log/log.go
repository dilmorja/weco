// go:build js && wasm
// +build js,wasm

package main

import "github.com/dilmorja/weco/window"

func main() {

	/// window.Console ///
	// TEST: Log
	window.Console.Log("window.Console.Log: TEST PASSED")

	// TEST: Info
	window.Console.Info("window.Console.Info: TEST PASSED")

	// TEST: Warn
	window.Console.Warn("window.Console.Warn: TEST PASSED")

	// TEST: Error
	window.Console.Error("window.Console.Error: TEST PASSED")
}