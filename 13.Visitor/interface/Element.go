package _interface

type Element interface {
	Accept(Visitor)
}
