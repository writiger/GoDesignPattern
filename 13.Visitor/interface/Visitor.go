package _interface

type Visitor interface {
	VisitFile(File)
	VisitDirectory(Directory)
}
