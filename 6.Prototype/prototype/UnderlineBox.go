package prototype

import (
	_interface "6.Prototype/interface"
	"fmt"
)

type UnderlinePen struct {
	decochar string
}

func InitUnderlinePen(s string) UnderlinePen {
	return UnderlinePen{decochar: s}
}

func (u UnderlinePen) Use(s string) {
	fmt.Println(`"`, s, `"`)
	length := len(s)
	for i := 0; i < length+4; i++ {
		fmt.Print(u.decochar)
	}
	fmt.Println()
}

func (u UnderlinePen) Clone() _interface.Product {
	return UnderlinePen{decochar: u.decochar}
}
