package tui

// CoreRoot is the common content container
type CoreRoot struct {
	*Component
}

// UIRoot to create core container
func UIRoot(o ...interface{}) *CoreRoot {
	c := &CoreRoot{Extends()}
	c.Children(o...)
	return c
}

// Render to render component content
func (v *CoreRoot) Render() interface{} {
	if o, ok := v.props.Get("children"); ok {
		if a, ok := o.([]interface{}); ok {
			if l := len(a); l == 1 {
				return a[0]
			}
		}
		return o
	}
	return make([]interface{}, 0)
}
