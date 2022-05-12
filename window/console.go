package window

import(
	. "syscall/js"
)

type ConsoleObject Value

// Get the native value
func (c ConsoleObject) AsValue() Value { return Value(c) }

// Writes an error message to the console if a assertion is false
func (c ConsoleObject) Assert(data ...interface{}) {
	c.AsValue().Call("assert", data...)
}

// Clears the console
func (c ConsoleObject) Clear() { c.AsValue().Call("clear") }

// Logs the number of times that this particular call to count() has been called
func (c ConsoleObject) Count(label ...interface{}) {
	if len(label) > 0 {
		c.AsValue().Call("count", label[0])
	} else {
		c.AsValue().Call("count")
	}
}

// Outputs an error message to the console
func (c ConsoleObject) Error(data ...interface{}) { c.AsValue().Call("error", data...) }

// Creates a new inline group in the console.
// This indents following console messages by an
// additional level, until window.Console.GroupEnd() is called
func (c ConsoleObject) Group(label ...interface{}) {
	if len(label) > 0 {
		c.AsValue().Call("group", label[0])
	} else {
		c.AsValue().Call("group")
	}
}

// Creates a new inline group in the console.
// However, the new group is created collapsed.
// The user will need to use the disclosure button to expand it
func (c ConsoleObject) GroupCollapsed(label ...interface{}) {
	if len(label) > 0 {
		c.AsValue().Call("groupCollapsed", label[0])
	} else {
		c.AsValue().Call("groupCollapsed")
	}
}

// Exits the current inline group in the console
func (c ConsoleObject) GroupEnd(label ...interface{}) {
	if len(label) > 0 {
		c.AsValue().Call("groupEnd", label[0])
	} else {
		c.AsValue().Call("groupEnd")
	}
}

// Outputs an informational message to the console
func (c ConsoleObject) Info(data ...interface{}) { c.AsValue().Call("info", data...) }

// Outputs a message to the console
func (c ConsoleObject) Log(data ...interface{}) { c.AsValue().Call("log", data...) }

// Displays tabular data as a table
func (c ConsoleObject) Table(tabularData interface{}, props ...interface{}) {
	if len(props) > 0 {
		c.AsValue().Call("table", tabularData, props[0])
	} else {
		c.AsValue().Call("table", tabularData)
	}
}

// Starts a timer (can track how long an operation takes)
func (c ConsoleObject) Time(label ...interface{}) {
	if len(label) > 0 {
		c.AsValue().Call("time", label[0])
	} else {
		c.AsValue().Call("time")
	}
}

// Stops a timer that was previously started by window.Console.Time()
func (c ConsoleObject) TimeEnd(label ...interface{}) {
	if len(label) > 0 {
		c.AsValue().Call("timeEnd", label[0])
	} else {
		c.AsValue().Call("timeEnd")
	}
}

// Outputs a stack trace to the console
func (c ConsoleObject) Trace(label ...interface{}) {
	if len(label) > 0 {
		c.AsValue().Call("trace", label[0])
	} else {
		c.AsValue().Call("trace")
	}
}

// Outputs a warning message to the console
func (c ConsoleObject) Warn(data ...interface{}) { c.AsValue().Call("warn", data...) }