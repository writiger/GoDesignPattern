package print

import (
	_interface "Adapter/interface"
	"fmt"
)

type Banner struct {
	_interface.Print
	Info string
}

func (b *Banner) ShowWithParen() {
	fmt.Println("(", b.Info, ")")
}

func (b *Banner) ShowWithAster() {
	fmt.Println("*", b.Info, "*")
}
