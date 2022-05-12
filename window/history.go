// Copyright (c) 2022, Daniel M. Jaén
// All rights reserved.

// https://www.w3schools.com/jsref/prop_win_history.asp
package window

import "syscall/js"

type HistoryObject js.Value

// Get the native value
func (h HistoryObject) AsValue() { return js.Value(h) }

// Loads the previous URL (page) in the history list
func (h HistoryObject) Back() { h.AsValue().Call("back") }

// Loads the next URL (page) in the history list
func (h HistoryObject) Forward() { h.AsValue().Call("forward") }

// Loads a specific URL (page) from the history list
func (h HistoryObject) Go(number int) { h.AsValue().Call("go", number) }

// Returns the number of URLs (pages) in the history list
func (h HistoryObject) Length() int {
	return h.AsValue().Get("length").Int()
}