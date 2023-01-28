package responsibility

import _interface "14.ChainOfResponsibility/interface"

type LimitSupport struct {
	name  string
	next  _interface.Support
	limit int
}

func InitLimitSupport(name string, limit int) *LimitSupport {
	return &LimitSupport{
		name:  name,
		next:  nil,
		limit: limit,
	}
}

func (ls *LimitSupport) SetNext(next _interface.Support) _interface.Support {
	ls.next = next
	return next
}

func (ls *LimitSupport) GetNext() _interface.Support {
	return ls.next
}

func (ls *LimitSupport) GetName() string {
	return ls.name
}

func (ls *LimitSupport) Resolve(t _interface.Trouble) bool {
	if t.GetNumber() >= ls.limit {
		return false
	}
	return true
}

func (ls *LimitSupport) Support(t _interface.Trouble) {
	SupportImpl.CommonResolve(ls, t)
}
