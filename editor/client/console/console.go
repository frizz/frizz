package console // import "kego.io/editor/client/console"

// ke: {"package": {"notest": true}}

import "github.com/gopherjs/gopherjs/js"

var c *js.Object

func initialise() {
	if c == nil {
		c = js.Global.Get("console")
	}
}

// Error is like Log but also includes a stack trace.
func Error(objs ...interface{}) {
	initialise()
	c.Call("error", objs...)
}

// Log displays a message in the console. You pass one or more objects
// to this method, each of which are evaluated and concatenated into a
// space-delimited string. The first parameter you pass to Log may
// contain format specifiers.
func Log(objs ...interface{}) {
	initialise()
	c.Call("log", objs...)
}
