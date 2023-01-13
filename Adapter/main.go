package main

import myPrint "Adapter/print"

func main() {
	p := myPrint.InitPrintBanner("Hello")
	p.PrintWeak()
	p.PrintStrong()
}
