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
)