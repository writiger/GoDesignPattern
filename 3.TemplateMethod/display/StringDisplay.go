package display

import "fmt"

type StringDisplay struct {
	Info string
}

type stringDisplayTemplate struct {
	Info string
	Len  int
}

func (d stringDisplayTemplate) Open() {
	d.Println()
}
func (d stringDisplayTemplate) Print() {
	fmt.Println("|" + d.Info + "|")
}
func (d stringDisplayTemplate) Close() {
	d.Println()
}
func (d stringDisplayTemplate) Println() {
	fmt.Print("+")
	for i := 0; i < d.Len; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

func (sd *StringDisplay) Display() {
	Display(stringDisplayTemplate{
		Info: sd.Info,
		Len:  len(sd.Info),
	})
}

func InitStringDisplay(info string) StringDisplay {
	return StringDisplay{Info: info}
}
