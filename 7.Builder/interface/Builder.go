package _interface

type Builder interface {
	MakeTitle(string)
	MakeString(string)
	MakeItems(...string)
	Close()
}
