package abstract_parts

import _interface "8.AbstractFactory/interface"

type Page interface {
	_interface.Item
	Add(_interface.Item)
	Output()
}
