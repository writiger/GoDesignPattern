package main

import (
	_interface "6.Prototype/interface"
	"6.Prototype/prototype"
)

func main() {
	manager := prototype.Manager{Showcase: map[string]_interface.Product{}}
	uPen := prototype.InitUnderlinePen("_")
	mBox := prototype.InitMessageBox("*")
	sBox := prototype.InitMessageBox("/")
	manager.Register("strong message", uPen)
	manager.Register("warning box", mBox)
	manager.Register("slash box", sBox)

	p1 := manager.Create("strong message")
	p1.Use("Hello World")
	p2 := manager.Create("warning box")
	p2.Use("Hello World")
	p3 := manager.Create("slash box")
	p3.Use("Hello World")
}
