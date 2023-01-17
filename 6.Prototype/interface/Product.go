package _interface

type Product interface {
	Use(string)
	Cloneable
}

type Cloneable interface {
	Clone() Product
}
