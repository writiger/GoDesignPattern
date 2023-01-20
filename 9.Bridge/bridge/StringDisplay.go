package bridge

import "fmt"

// StringDisplay 新增实现
type StringDisplay struct {
	info  string
	width int
}

func InitStringDisplay(info string) *StringDisplay {
	return &StringDisplay{
		info:  info,
		width: len(info),
	}
}

func (sd StringDisplay) RawOpen() {
	sd.printLine()
}

func (sd StringDisplay) RawPrint() {
	fmt.Println("|" + sd.info + "|")
}

func (sd StringDisplay) RawClose() {
	sd.printLine()
}

func (sd StringDisplay) printLine() {
	fmt.Print("+")
	for i := 0; i < sd.width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}
