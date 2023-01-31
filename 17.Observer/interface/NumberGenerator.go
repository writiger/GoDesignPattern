package _interface

type NumberGenerator interface {
	AddObserver(Observer)
	DeleteObserver(Observer)
	NotifyObservers()
	GetNumber() int
	Execute()
}
