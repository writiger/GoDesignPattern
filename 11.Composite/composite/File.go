package composite

import (
	"fmt"
)

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

func (f *File) GetName() string {
	return f.name
}

func (f *File) GetSize() int {
	return f.size
}

func (f *File) PrintList(prefix string) {
	fmt.Println(prefix, "/", EntryToString(f))
}
