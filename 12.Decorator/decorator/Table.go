package decorator

import (
	_interface "12.Decorator/interface"
	"fmt"
)

type Table struct {
}

func (t Table) Show(display _interface.Display) {
	rows := display.GetRows()
	for i := 0; i < rows; i++ {
		fmt.Println(display.GetRowText(i))
	}
}
