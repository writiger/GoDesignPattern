package main

import (
	"7.Builder/builder"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: 7.builder 
example: go run main.go plain`)
		os.Exit(1)
	}
	way := os.Args[1]
	switch way {
	case "plain":
		fmt.Println(makePlain())
	case "html":
		fmt.Println(makeHtml())
	default:
		fmt.Fprintf(os.Stderr, `usage: 7.builder 
methods: plain / html`)
		os.Exit(1)
	}
}

func makePlain() string {
	tb := builder.InitTextBuilder()
	d := builder.InitDirector(tb)
	d.Construct()
	return tb.GetResult()
}
func makeHtml() string {
	hb := builder.InitHtmlBuilder("Greeting")
	d := builder.InitDirector(hb)
	d.Construct()
	return "文件" + hb.GetResult() + "编写完成\n"
}
