package tui

// CoreRoot is the common content container
type CoreRoot struct {
	*Component
}

// UIRoot to create core container
func UIRoot(o ...interface{}) *CoreRoot {
	c := &CoreRoot{Extends()}
	c.Children(o...)
	c.Style().Flex = 1
	c.Style().Padding = 2
	return c
}

// Render to render component content
func (v *CoreRoot) Render() interface{} {
	return uilayer(Attr{
		Style: v.style,
		Props: v.props,
	}, v.childrens())
}
