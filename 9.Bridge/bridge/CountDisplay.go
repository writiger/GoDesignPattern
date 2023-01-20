package bridge

import _interface "9.Bridge/interface"

// CountDisplay 新增功能
type CountDisplay struct {
	Display
}

func InitCountDisplay(impl _interface.DisplayImpl) *CountDisplay {
	return &CountDisplay{Display: *InitDisplay(impl)}
}

func (cd CountDisplay) MultiDisplay(times int) {
	cd.open()
	for i := 0; i < times; i++ {
		cd.print()
	}
	cd.close()
}
