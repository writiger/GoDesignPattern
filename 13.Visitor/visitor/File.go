package visitor

import _interface "13.Visitor/interface"

type File struct {
	name string
	size int
}

func InitFile(name string, size int) *File {
	return &File{
		name: name,
		size: size,
	}
}

func (f File) GetName() string {
	return f.name
}

func (f File) GetSize() int {
	return f.size
}

func (f File) Accept(visitor _interface.Visitor) {
	visitor.VisitFile(f)
}
