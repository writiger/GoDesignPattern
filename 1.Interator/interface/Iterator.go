package _interface

type Iterator interface {
	HasNext() bool
	Next() interface{}
}
