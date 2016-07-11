package tui

// Change is the component state object
type Change map[string]interface{}

// Set to set data from state map
func (v Change) Set(s string, i interface{}) {
	v[s] = i
}

// Get to retrieve data from state map
func (v Change) Get(s string) interface{} {
	return v[s]
}

// Keys returns list of variables keys inside the Change object
func (v Change) Keys() []string {
	var i int
	k := make([]string, len(v))
	for s := range v {
		k[i] = s
		i++
	}
	return k
}
