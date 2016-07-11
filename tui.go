package tui

// Render to render component in terminal
func Render(i interface{}) []interface{} {
	tmpl := compile(i)
	return tmpl
}

// Render to render component in terminal
func compile(i interface{}) []interface{} {
	tmpl := build(i)
	if len(tmpl) == 1 {
		if o, ok := tmpl[0].([]interface{}); ok {
			tmpl = o
		}
	}
	return tmpl
}

func build(i interface{}) []interface{} {
	parent := []interface{}{}
	switch o := i.(type) {
	case IComponent:
		tmpl := o.Render()
		if tmpl == nil {
			return parent
		}
		temp := build(tmpl)
		if l := len(temp); l > 1 {
			parent = append(parent, temp)
		} else if l == 1 {
			parent = append(parent, temp[0])
		}
	case []interface{}:
		sub := []interface{}{}
		for _, v := range o {
			if v == nil {
				continue
			}
			temp := build(v)
			if l := len(temp); l > 1 {
				sub = append(sub, temp)
			} else if l == 1 {
				sub = append(sub, temp[0])
			}
		}
		parent = append(parent, sub)
	default:
		parent = append(parent, o)
	}
	return parent
}
