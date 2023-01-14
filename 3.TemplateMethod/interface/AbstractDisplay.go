package _interface

type AbstractDisplay interface {
	Open()
	Print()
	Close()
}

type Display interface {
	Display()
}
