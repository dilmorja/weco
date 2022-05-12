// Copyright (c) 2022, Daniel M. Jaén
// All rights reserved.

// https://www.w3schools.com/jsref/obj_window.asp
package window

import(
	. "syscall/js"
)

var w = Global().Get("window")

var(
	// Returns a boolean true if a window is closed.
	Closed = w.Get("closed").Bool()
	// Returns the Console Object for the window.
	// See also https://www.w3schools.com/jsref/obj_console.asp
	Console = ConsoleObject(w.Get("console"))
	// Returns the Document object for the window.
	// See also https://www.w3schools.com/jsref/dom_obj_document.asp
	Document = w.Get("document")
	// Returns the frame in which the window runs.
	FrameElement = w.Get("frameElement")
	// Returns all window objects running in the window.
	Frames = w.Get("frames")
	// Returns the History object for the window.
	// See also https://www.w3schools.com/jsref/obj_history.asp
	History = HistoryObject(w.Get("history"))
)