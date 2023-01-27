package main

import (
	"13.Visitor/visitor"
	"fmt"
)

func main() {
	fmt.Println("Making root entries...")
	rootDir := visitor.InitDirectory("root")
	binDir := visitor.InitDirectory("bin")
	tmpDir := visitor.InitDirectory("tmp")
	usrDir := visitor.InitDirectory("usr")
	rootDir.Add(binDir)
	rootDir.Add(tmpDir)
	rootDir.Add(usrDir)
	binDir.Add(visitor.InitFile("vi", 10000))
	binDir.Add(visitor.InitFile("latex", 20000))
	rootDir.Accept(visitor.InitVisitor())

	fmt.Printf("\nMaking usr entries...\n")
	zsDir := visitor.InitDirectory("zhangsan")
	lsDir := visitor.InitDirectory("lisi")
	wwDir := visitor.InitDirectory("wangwu")
	usrDir.Add(zsDir)
	usrDir.Add(lsDir)
	usrDir.Add(wwDir)
	zsDir.Add(visitor.InitFile("diary.html", 100))
	zsDir.Add(visitor.InitFile("main.go", 200))
	lsDir.Add(visitor.InitFile("memo.tex", 300))
	wwDir.Add(visitor.InitFile("game.doc", 400))
	wwDir.Add(visitor.InitFile("junk.mail", 500))
	rootDir.Accept(visitor.InitVisitor())
}
