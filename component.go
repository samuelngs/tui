package tui

// IComponent is the terminal component interface
type IComponent interface {
	// public
	Ref() string
	Render() interface{}
	State() *Map
	Props() *Map
	Style() *Style
	SetState(*Change)
	ComponentWillUpdate(*Map, *Map)
	ComponentDidUpdate(*Map, *Map)
	ShouldComponentUpdate(*Map, *Map) bool
	Children(...interface{})
	// private
	render() interface{}
	redraw()
	area() *Bound
	childrens() interface{}
}

// Extends to extend custom component with core component
func Extends(m ...*Map) *Component {
	var props *Map
	for _, s := range m {
		props = s
		break
	}
	if props == nil {
		props = &Map{m: make(map[string]interface{})}
	}
	if props.m == nil {
		props.m = make(map[string]interface{})
	}
	if _, ok := props.Get("children"); !ok {
		props.Set("children", make([]interface{}, 0))
	}
	return &Component{
		id:    random(),
		style: new(Style),
		state: &Map{m: make(map[string]interface{})},
		props: props,
		bound: new(Style).GetWindowSize(),
	}
}

// Attr the attribute struct
type Attr struct {
	Style *Style
	Props *Map
}

// Component struct
type Component struct {
	id    string
	state *Map
	props *Map
	style *Style
	bound *Bound
}

// State returns component state
func (v *Component) State() *Map {
	return v.state
}

// Props returns passed in props
func (v *Component) Props() *Map {
	return v.props
}

// SetState to set data to component state
func (v *Component) SetState(c *Change) {
	ks := c.Keys()
	for _, k := range ks {
		o := c.Get(k)
		v.State().Set(k, o)
	}
	go v.redraw()
}

// Ref to return component id
func (v *Component) Ref() string {
	return v.id
}

// Render ui handler
func (v *Component) Render() interface{} {
	return nil
}

// ComponentWillUpdate - Invoked immediately before rendering when new props or state are being received.
func (v *Component) ComponentWillUpdate(nextProps *Map, nextState *Map) {}

// ComponentDidUpdate - Invoked immediately after the component's updates are flushed to the terminal.
func (v *Component) ComponentDidUpdate(prevProps *Map, prevState *Map) {}

// ShouldComponentUpdate - Invoked before rendering when new props or state are being received
func (v *Component) ShouldComponentUpdate(nextProps *Map, nextState *Map) bool {
	return true
}

// Children to add component children
func (v *Component) Children(i ...interface{}) {
	var chrn []interface{}
	o, ok := v.props.Get("children")
	if !ok {
		v.props.Set("children", make([]interface{}, 0))
	}
	if _, ok := o.([]interface{}); !ok {
		v.props.Set("children", make([]interface{}, 0))
	}
	o, _ = v.props.Get("children")
	chrn = o.([]interface{})
	for _, c := range i {
		switch s := c.(type) {
		case Attr:
			if s.Props != nil {
				v.props = s.Props
			}
			if s.Style != nil {
				v.style = s.Style
			}
		case *Attr:
			if s.Props != nil {
				v.props = s.Props
			}
			if s.Style != nil {
				v.style = s.Style
			}
		case Style:
			v.style = &s
		case *Style:
			v.style = s
		default:
			chrn = append(chrn, c)
		}
	}
	v.props.Set("children", chrn)
}

// Style to get style from component
func (v *Component) Style() *Style {
	return v.style
}

// (re)-draw to (re)draw user interface
func (v *Component) redraw() {
	v.Render()
}

// private render handler
func (v *Component) render() interface{} {
	return v.Render()
}

// resize handler
func (v *Component) resize() {
	v.Render()
}

// area returns the availabe bound size
func (v *Component) area() *Bound {
	return v.bound
}

// area returns the availabe bound size
func (v *Component) childrens() interface{} {
	var children interface{}
	if o, ok := v.props.Get("children"); ok {
		if a, ok := o.([]interface{}); ok {
			if l := len(a); l == 1 {
				children = a[0]
			} else {
				children = a
			}
		}
	}
	return children
}
