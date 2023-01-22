package composite

import (
	_interface "11.Composite/interface"
	"fmt"
)

type Directory struct {
	name      string
	directory []*_interface.Entry
}

func InitDirectory(name string) *Directory {
	return &Directory{
		name:      name,
		directory: make([]*_interface.Entry, 0),
	}
}

func (d *Directory) Add(entry _interface.Entry) {
	d.directory = append(d.directory, &entry)
}

func (d *Directory) GetName() string {
	return d.name
}

func (d *Directory) GetSize() int {
	res := 0
	for _, item := range d.directory {
		res += (*item).GetSize()
	}
	return res
}

func (d *Directory) PrintList(prefix string) {
	fmt.Println(prefix, "/", EntryToString(d))
	for _, item := range d.directory {
		(*item).PrintList(prefix + " / " + d.name)
	}
}
