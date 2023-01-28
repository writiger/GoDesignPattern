package responsibility

import _interface "14.ChainOfResponsibility/interface"

type NoSupport struct {
	name string
	next _interface.Support
}

func InitNoSupport(name string) *NoSupport {
	return &NoSupport{
		name: name,
		next: nil,
	}
}

func (ns *NoSupport) SetNext(next _interface.Support) _interface.Support {
	ns.next = next
	return next
}

func (ns *NoSupport) GetNext() _interface.Support {
	return ns.next
}

func (ns *NoSupport) GetName() string {
	return ns.name
}

func (ns *NoSupport) Resolve(_interface.Trouble) bool {
	return false
}

func (ns *NoSupport) Support(t _interface.Trouble) {
	SupportImpl.CommonResolve(ns, t)
}
