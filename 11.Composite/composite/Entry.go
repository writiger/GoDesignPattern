package composite

import (
	_interface "11.Composite/interface"
	"strconv"
)

func EntryToString(e _interface.Entry) string {
	return e.GetName() + " (" + strconv.Itoa(e.GetSize()) + ")"
}
