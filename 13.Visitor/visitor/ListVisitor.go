package visitor

import (
	_interface "13.Visitor/interface"
	"fmt"
	"strconv"
)

type Visitor struct {
	currentDir string
}

func InitVisitor() *Visitor {
	return &Visitor{currentDir: ""}
}

func (v *Visitor) VisitFile(f _interface.File) {
	fmt.Println(v.currentDir + "/" + entryToString(f))
}
func (v *Visitor) VisitDirectory(d _interface.Directory) {
	fmt.Println(v.currentDir + "/" + entryToString(d))
	saveDir := v.currentDir
	v.currentDir = v.currentDir + "/" + d.GetName()
	for _, item := range d.GetRange() {
		item.Accept(v)
	}
	v.currentDir = saveDir
}

func entryToString(e _interface.Entry) string {
	return e.GetName() + " (" + strconv.Itoa(e.GetSize()) + ")"
}
