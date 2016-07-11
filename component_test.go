package tui

import (
	"fmt"
	"testing"
)

type testComponent1 struct {
	*Component
}

func (v *testComponent1) Render() interface{} {
	return v.Props().Get("")
}

type testComponent2 struct {
	*Component
}

func (v *testComponent2) Render() interface{} {
	return nil
}

func TestCreateComponent(t *testing.T) {
	c := &testComponent1{Extends()}
	o := c.Props()
	fmt.Println(c, o)
}
