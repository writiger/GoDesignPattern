package main

import factory2 "4.FactoryMethod/factory"

func main() {
	factory := factory2.InitIDCardFactory()
	card1 := factory.Create("小明")
	card2 := factory.Create("小红")
	card3 := factory.Create("小钢")
	card1.Use()
	card2.Use()
	card3.Use()
}
