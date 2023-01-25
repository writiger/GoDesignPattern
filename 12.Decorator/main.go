package main

import (
	"12.Decorator/decorator"
)

func main() {
	t := &decorator.Table{}
	b1 := decorator.InitStringDisplay("Hello,World!")
	b2 := decorator.InitSideBorder(b1, "#")
	b3 := decorator.InitFullBorder(b2)
	t.Show(b1)
	t.Show(b2)
	t.Show(b3)
	b4 := decorator.InitSideBorder(
		decorator.InitFullBorder(
			decorator.InitFullBorder(
				decorator.InitSideBorder(
					decorator.InitFullBorder(
						decorator.InitStringDisplay("Hello,World.")),
					"*",
				)),
		),
		"/",
	)
	t.Show(b4)
}
