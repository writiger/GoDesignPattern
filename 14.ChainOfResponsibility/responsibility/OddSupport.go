package responsibility

import _interface "14.ChainOfResponsibility/interface"

type OddSupport struct {
	name string
	next _interface.Support
}

func InitOddSupport(name string) *OddSupport {
	return &OddSupport{
		name: name,
		next: nil,
	}
}

func (os *OddSupport) SetNext(next _interface.Support) _interface.Support {
	os.next = next
	return next
}

func (os *OddSupport) GetNext() _interface.Support {
	return os.next
}

func (os *OddSupport) GetName() string {
	return os.name
}

func (os *OddSupport) Resolve(t _interface.Trouble) bool {
	if t.GetNumber()%2 == 0 {
		return false
	}
	return true
}

func (os *OddSupport) Support(t _interface.Trouble) {
	SupportImpl.CommonResolve(os, t)
}
