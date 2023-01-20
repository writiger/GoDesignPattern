package bridge

import _interface "9.Bridge/interface"

type Display struct {
	impl _interface.DisplayImpl
}

func InitDisplay(d _interface.DisplayImpl) *Display {
	return &Display{d}
}

func (d Display) open() {
	d.impl.RawOpen()
}

func (d Display) print() {
	d.impl.RawPrint()
}

func (d Display) close() {
	d.impl.RawClose()
}

func (d Display) Display() {
	d.open()
	d.print()
	d.close()
}
