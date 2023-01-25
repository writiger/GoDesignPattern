package decorator

import (
	_interface "12.Decorator/interface"
	"strings"
)

type FullBorder struct {
	display _interface.Display
}

func InitFullBorder(display _interface.Display) *FullBorder {
	return &FullBorder{display: display}
}

func (fb FullBorder) GetColumns() int {
	return 1 + fb.display.GetColumns() + 1
}

func (fb FullBorder) GetRows() int {
	return 1 + fb.display.GetRows() + 1
}

func (fb FullBorder) GetRowText(row int) string {
	res := strings.Builder{}
	if row == 0 || row == fb.display.GetRows()+1 {
		res.WriteString("+")
		res.WriteString(makeLine("-", fb.display.GetColumns()))
		res.WriteString("+")
	} else {
		res.WriteString("|")
		res.WriteString(fb.display.GetRowText(row - 1))
		res.WriteString("|")
	}
	return res.String()
}

func makeLine(char string, count int) string {
	buf := strings.Builder{}
	for i := 0; i < count; i++ {
		buf.WriteString(char)
	}
	return buf.String()
}
