package tui

import "bytes"

// view is the common content container
type view struct {
	*Component
}

// uicore to create core container
func uilayer(o ...interface{}) *view {
	c := &view{Extends()}
	c.Children(o...)
	return c
}

// Render to render component content
func (v *view) Render() interface{} {
	b := new(bytes.Buffer)
	if o, ok := v.props.Get("children"); ok {
		if a, ok := o.([]interface{}); ok {
			if l := len(a); l == 1 {
				b.Write(v.render(a[0]))
			} else if l > 0 {
				for _, c := range a {
					b.Write(v.render(c))
				}
			}
		}
	}
	return b.String()
}

func (v *view) render(i interface{}) []byte {
	area := v.area()
	b := new(bytes.Buffer)
	for i := 0; i < int(area.Height); i++ {
		for y := 0; y < int(area.Width); y++ {
			b.WriteRune('.')
			if y == int(area.Width)-1 {
				b.WriteRune('\n')
			}
		}
	}
	switch {
	case v.Style().BorderRightWidth > 0:
		// v.renderBorderRight(b)
	}
	return b.Bytes()
}

// func (v *view) renderBorderRight(b []byte) {
// 	area := v.area()
// 	for i := 0; i < int(area.Height); i++ {
// 	}
// }
