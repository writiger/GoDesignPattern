package main

import (
	"11.Composite/composite"
	"fmt"
)

func main() {
	fmt.Println("Making root entries...")
	rootdir := composite.InitDirectory("root")
	bindir := composite.InitDirectory("bin")
	tmpdir := composite.InitDirectory("tmp")
	usrdir := composite.InitDirectory("usr")
	rootdir.Add(bindir)
	rootdir.Add(tmpdir)
	rootdir.Add(usrdir)
	bindir.Add(composite.InitFile("vi", 10000))
	bindir.Add(composite.InitFile("latex", 20000))
	fmt.Printf("\nMaking usr entries...\n")
	zsdir := composite.InitDirectory("zhangsan")
	lsdir := composite.InitDirectory("lisi")
	wwdir := composite.InitDirectory("wangwu")
	usrdir.Add(zsdir)
	usrdir.Add(lsdir)
	usrdir.Add(wwdir)
	zsdir.Add(composite.InitFile("diary.html", 100))
	zsdir.Add(composite.InitFile("main.go", 200))
	lsdir.Add(composite.InitFile("memo.tex", 300))
	wwdir.Add(composite.InitFile("game.doc", 400))
	wwdir.Add(composite.InitFile("junk.mail", 500))
	rootdir.PrintList("")
}
