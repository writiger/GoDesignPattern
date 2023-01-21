package _interface

type Hand interface {
	IsStrongerThan(Hand) bool
	IsWeakerThan(Hand) bool
	GetHandValue() int
}
