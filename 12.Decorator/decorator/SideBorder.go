package decorator

import _interface "12.Decorator/interface"

type SideBorder struct {
	display _interface.Display
	char    string
}

func InitSideBorder(display _interface.Display, char string) *SideBorder {
	return &SideBorder{
		display: display,
		char:    char,
	}
}

func (sb SideBorder) GetColumns() int {
	return 1 + sb.display.GetColumns() + 1
}

func (sb SideBorder) GetRows() int {
	return sb.display.GetRows()
}

func (sb SideBorder) GetRowText(row int) string {
	return sb.char + sb.display.GetRowText(row) + sb.char
}
