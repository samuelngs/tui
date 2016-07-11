package tui

// CoreView is the common content container
type CoreView struct {
	*Component
}

// UIView to create core container
func UIView(o ...interface{}) *CoreView {
	c := &CoreView{Extends()}
	c.Children(o...)
	return c
}

// Render to render component content
func (v *CoreView) Render() interface{} {
	if o, ok := v.props.Get("children"); ok {
		return o
	}
	return make([]interface{}, 0)
}
