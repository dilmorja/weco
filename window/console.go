package window

import(
	. "syscall/js"
)

type ConsoleObject Value

func (c ConsoleObject) Assert(condition bool, data ...interface{}) {
	c.Call("assert", condition, data...)
}

func (c ConsoleObject) Clear() { c.Call("clear") }

func (c ConsoleObject) Count(label ...interface{}) {
	if len(label) > 0 {
		c.Call("count", label[0])
	} else {
		c.Call("count")
	}
}

func (c ConsoleObject) Error(data ...interface{}) { c.Call("error", data...) }

func (c ConsoleObject) Group(label ...interface{}) {
	if len(label) > 0 {
		c.Call("group", label[0])
	} else {
		c.Call("group")
	}
}

func (c ConsoleObject) GroupCollapsed(label ...interface{}) {
	if len(label) > 0 {
		c.Call("groupCollapsed", label[0])
	} else {
		c.Call("groupCollapsed")
	}
}

func (c ConsoleObject) GroupEnd(label ...interface{}) {
	if len(label) > 0 {
		c.Call("groupEnd", label[0])
	} else {
		c.Call("groupEnd")
	}
}

func (c ConsoleObject) Info(data ...interface{}) { c.Call("info", data...) }

func (c ConsoleObject) Log(data ...interface{}) { c.Call("log", data...) }

func (c ConsoleObject) Table(tabularData interface{}, props ...interface{}) {
	if len(props) > 0 {
		c.Call("table", tabularData, props[0])
	} else {
		c.Call("table", tabularData)
	}
}

func (c ConsoleObject) Time(label ...interface{}) {
	if len(label) > 0 {
		c.Call("time", label[0])
	} else {
		c.Call("time")
	}
}

func (c ConsoleObject) TimeLog(label ...interface{}, data ...interface{}) {
	if len(label) > 0 {
		if len(data) > 0 {
			c.Call("timeLog", label[0], data...)
		} else {
			c.Call("timeLog", label[0])
		}
	} else {
		c.Call("timeLog")
	}
}

func (c ConsoleObject) TimeEnd(label ...interface{}) {
	if len(label) > 0 {
		c.Call("timeEnd", label[0])
	} else {
		c.Call("timeEnd")
	}
}

func (c ConsoleObject) Trace(label ...interface{}) {
	if len(label) > 0 {
		c.Call("trace", label[0])
	} else {
		c.Call("trace")
	}
}

func (c ConsoleObject) Warn(data ...interface{}) { c.Call("warn", data...) }