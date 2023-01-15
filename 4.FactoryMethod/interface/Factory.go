package _interface

import "4.FactoryMethod/factory"

type Factory interface {
	CreateProduct(string) factory.IDCard
	RegisterProduct(string)
	Create(string) factory.IDCard
}
