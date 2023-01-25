package _interface

type Display interface {
	GetColumns() int
	GetRows() int
	GetRowText(int) string
}

type Show interface {
	Show(Display)
}
