package _interface

type Directory interface {
	Entry
	Add(Entry)
	GetRange() []Entry
}
