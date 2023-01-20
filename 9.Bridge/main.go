package main

import "9.Bridge/bridge"

func main() { 
	sd1 := bridge.InitStringDisplay("Hello,China.")
	sd2 := bridge.InitStringDisplay("Hello,World.")
	sd3 := bridge.InitStringDisplay("Hello,Universe.")
	d1 := bridge.InitDisplay(sd1)
	d2 := bridge.InitCountDisplay(sd2)
	d3 := bridge.InitCountDisplay(sd3)

	d1.Display()
	d2.Display.Display()
	d3.MultiDisplay(5)
}
