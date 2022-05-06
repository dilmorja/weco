package window

import(
	. "syscall/js"
)

type ConsoleObject Value

func (c ConsoleObject) AsValue() Value { return Value(c) }

func (c ConsoleObject) Assert(data ...interface{}) {
	c.AsValue().Call("assert", data...)
}

func (c ConsoleObject) Clear() { c.AsValue().Call("clear") }

func (c ConsoleObject) Count(label ...interface{}) {
	if len(label) > 0 {
		c.AsValue().Call("count", label[0])
	} else {
		c.AsValue().Call("count")
	}
}

func (c ConsoleObject) Error(data ...interface{}) { c.AsValue().Call("error", data...) }

func (c ConsoleObject) Group(label ...interface{}) {
	if len(label) > 0 {
		c.AsValue().Call("group", label[0])
	} else {
		c.AsValue().Call("group")
	}
}

func (c ConsoleObject) GroupCollapsed(label ...interface{}) {
	if len(label) > 0 {
		c.AsValue().Call("groupCollapsed", label[0])
	} else {
		c.AsValue().Call("groupCollapsed")
	}
}

func (c ConsoleObject) GroupEnd(label ...interface{}) {
	if len(label) > 0 {
		c.AsValue().Call("groupEnd", label[0])
	} else {
		c.AsValue().Call("groupEnd")
	}
}

func (c ConsoleObject) Info(data ...interface{}) { c.AsValue().Call("info", data...) }

func (c ConsoleObject) Log(data ...interface{}) { c.AsValue().Call("log", data...) }

func (c ConsoleObject) Table(tabularData interface{}, props ...interface{}) {
	if len(props) > 0 {
		c.AsValue().Call("table", tabularData, props[0])
	} else {
		c.AsValue().Call("table", tabularData)
	}
}

func (c ConsoleObject) Time(label ...interface{}) {
	if len(label) > 0 {
		c.AsValue().Call("time", label[0])
	} else {
		c.AsValue().Call("time")
	}
}

func (c ConsoleObject) TimeEnd(label ...interface{}) {
	if len(label) > 0 {
		c.AsValue().Call("timeEnd", label[0])
	} else {
		c.AsValue().Call("timeEnd")
	}
}

func (c ConsoleObject) Trace(label ...interface{}) {
	if len(label) > 0 {
		c.AsValue().Call("trace", label[0])
	} else {
		c.AsValue().Call("trace")
	}
}

func (c ConsoleObject) Warn(data ...interface{}) { c.AsValue().Call("warn", data...) }