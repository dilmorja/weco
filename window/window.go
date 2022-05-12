// Copyright (c) 2022, Daniel M. Jaén
// All rights reserved.

// https://www.w3schools.com/jsref/obj_window.asp
package window

import(
	. "syscall/js"
)

var w = Global().Get("window")

var(
	Closed = w.Get("closed").Bool()
	Console = ConsoleObject(w.Get("console"))
	FrameElement = w.Get("frameElement")
	Frames = w.Get("frames")
	// https://www.w3schools.com/jsref/prop_win_document.asp
	Document = w.Get("document")
	// https://www.w3schools.com/jsref/obj_history.asp
	History = HistoryObject(w.Get("history"))
)