package display

import (
	"fmt"
)

type CharDisplay struct {
	Info string
}

type charDisplayTemplate struct {
	Info string
}

func (d charDisplayTemplate) Open() {
	fmt.Print("<<")
}
func (d charDisplayTemplate) Print() {
	fmt.Print(d.Info)
}
func (d charDisplayTemplate) Close() {
	fmt.Println(">>")
}

func (cd *CharDisplay) Display() {
	Display(charDisplayTemplate{Info: cd.Info})
}

func InitCharDisplay(info string) CharDisplay {
	return CharDisplay{Info: info}
}
