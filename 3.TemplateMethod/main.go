package main

import "TemplateMethod/display"

func main() {
	cd := display.InitCharDisplay("H")
	cd.Display()
	sd := display.InitStringDisplay("Hello,World!")
	sd.Display()
}
