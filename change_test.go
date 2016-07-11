package tui

import "testing"

func TestChangeSet(t *testing.T) {
	s := "test3"
	v := Change{}
	v.Set("name3", s)
	if o := v.Get("name3").(string); o != s {
		t.Errorf("expected %v, got %v", s, o)
	}
}

func TestChangeGet(t *testing.T) {
	s := -1
	v := Change{"name1": s}
	if o := v.Get("name1").(int); o != s {
		t.Errorf("expected %v, got %v", s, o)
	}
}

func TestChangeKeys(t *testing.T) {
	v := Change{"name1": 1, "name2": 2, "name3": 3}
	if s := len(v.Keys()); s != 3 {
		t.Errorf("expected size 3, got %v", s)
	}
}
