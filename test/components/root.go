package components

import "github.com/samuelngs/tui"

// HelloWorldComponent - the main controller component
type HelloWorldComponent struct {
	*tui.Component
}

// HelloWorldView to create root container
func HelloWorldView(o ...interface{}) *HelloWorldComponent {
	c := &HelloWorldComponent{tui.Extends()}
	c.Children(o...)
	return c
}

// Render component content
func (v *HelloWorldComponent) Render() interface{} {
	if o, ok := v.Props().Get("children"); ok {
		return o
	}
	return make([]interface{}, 0)
}
