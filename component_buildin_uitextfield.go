package tui

// CoreTextField is the common content container
type CoreTextField struct {
	*Component
}

// UITextField to create core container
func UITextField(o ...interface{}) *CoreTextField {
	c := &CoreTextField{Extends()}
	c.Children(o...)
	return c
}

// Render to render component content
func (v *CoreTextField) Render() interface{} {
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
