package main

import "github.com/samuelngs/tui"

func main() {

	t := tui.Term()

	c := tui.UIRoot(
	// tui.UIView(tui.Attr{
	// 	Style: &tui.Style{
	// 		Flex: 1,
	// 	},
	// }, "a"),
	// tui.UIView(tui.Attr{
	// 	Style: &tui.Style{
	// 		Flex: 2,
	// 	},
	// }, "b"),
	// components.HelloWorldView(
	// 	tui.UIView("c"),
	// 	tui.UIView(
	// 		"e",
	// 		tui.UIView("f",
	// 			tui.UIView("g", "h")),
	// 	),
	// ),
	// tui.UITextField(),
	)

	// t.Printf("1")
	// t.Flush()
	//
	// t.Printf("2")
	// t.Reset()
	// t.Flush()
	//
	// time.Sleep(2 * time.Second)
	//
	// t.Clear()
	// t.Flush()
	//
	// t.Printf("3")
	// t.Flush()

	t.Render(c)

	// fmt.Println(tui.Render(c), t.Size())
}
