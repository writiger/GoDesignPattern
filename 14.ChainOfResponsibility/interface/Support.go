package _interface

type Support interface {
	SetNext(Support) Support
	GetNext() Support
	GetName() string
	Resolve(Trouble) bool
	Support(Trouble)
}
