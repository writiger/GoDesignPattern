package prototype

import (
	_interface "6.Prototype/interface"
	"fmt"
)

type MessageBox struct {
	decochar string
}

func InitMessageBox(s string) MessageBox {
	return MessageBox{decochar: s}
}

func (m MessageBox) Use(s string) {
	length := len(s)
	for i := 0; i < length+4; i++ {
		fmt.Print(m.decochar)
	}
	fmt.Println()
	fmt.Println(m.decochar, s, m.decochar)
	for i := 0; i < length+4; i++ {
		fmt.Print(m.decochar)
	}
	fmt.Println()
}

func (m MessageBox) Clone() _interface.Product {
	return MessageBox{decochar: m.decochar}
}
