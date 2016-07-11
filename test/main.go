package main

import (
	"fmt"

	"github.com/samuelngs/tui"
	"github.com/samuelngs/tui/test/components"
)

func main() {

	c := tui.UIRoot(
		tui.UIView(tui.Attr{
			Style: &tui.Style{},
		}, "a"),
		tui.UIView("b"),
		components.HelloWorldView(
			tui.UIView("c"),
			tui.UIView(
				"e",
				tui.UIView("f"),
			),
		),
	)

	t := tui.Render(c)

	fmt.Println(t)
}
