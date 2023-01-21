package _interface

type Strategy interface {
	NextHand() Hand
	Study(bool)
}
