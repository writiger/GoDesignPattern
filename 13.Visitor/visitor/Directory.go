package visitor

import _interface "13.Visitor/interface"

type Directory struct {
	name string
	dir  []_interface.Entry
}

func InitDirectory(name string) *Directory {
	return &Directory{
		name: name,
		dir:  make([]_interface.Entry, 0),
	}
}

func (d *Directory) GetRange() []_interface.Entry {
	return d.dir
}

func (d *Directory) GetName() string {
	return d.name
}

func (d *Directory) GetSize() int {
	res := 0
	for _, item := range d.dir {
		res += item.GetSize()
	}
	return res
}

func (d *Directory) Add(entry _interface.Entry) {
	d.dir = append(d.dir, entry)
}

func (d *Directory) Accept(visitor _interface.Visitor) {
	visitor.VisitDirectory(d)
}
