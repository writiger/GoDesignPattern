package books

type BookNode struct {
	Next *BookNode
	Val  Book
}

func InitBookNode(book Book) *BookNode {
	return &BookNode{
		Next: nil,
		Val:  book,
	}
}
