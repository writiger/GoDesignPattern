package responsibility

import (
	_interface "14.ChainOfResponsibility/interface"
	"fmt"
)

type Support struct {
}

var SupportImpl *Support

func init() {
	SupportImpl = &Support{}
}

func (Support) CommonResolve(support _interface.Support, trouble _interface.Trouble) {
	if support.Resolve(trouble) {
		done(support.GetName(), trouble.ToString())
	} else if next := support.GetNext(); next != nil {
		next.Support(trouble)
	} else {
		fail(trouble.ToString())
	}
}

func done(s, t string) {
	fmt.Println(t, "is resolved by [", s, "]")
}

func fail(t string) {
	fmt.Println(t, "cannot be resolved.")
}
