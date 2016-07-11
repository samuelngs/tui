package tui

import "sync"

// Map is the component object
type Map struct {
	sync.RWMutex
	m map[string]interface{}
}

// Set to set data from map
func (v *Map) Set(s string, i interface{}) {
	v.Lock()
	v.m[s] = i
	v.Unlock()
}

// Get to retrieve data from map
func (v *Map) Get(s string) (interface{}, bool) {
	v.RLock()
	defer v.RUnlock()
	if o, ok := v.m[s]; ok {
		return o, ok
	}
	return nil, false
}

// Keys returns list of variables keys inside the map
func (v *Map) Keys() []string {
	var i int
	v.RLock()
	defer v.RUnlock()
	k := make([]string, len(v.m))
	for s := range v.m {
		k[i] = s
		i++
	}
	return k
}
