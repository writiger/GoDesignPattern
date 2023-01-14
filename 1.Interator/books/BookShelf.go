package books

type BookShelf struct {
	// 使用哑节点便于操作
	// 获取头节点时返回DummyHead即可
	DummyHead *BookNode
	Last      *BookNode
}

func InitBookShell() BookShelf {
	// 初始化一个BookNode作为哑节点
	bn := &BookNode{
		Next: nil,
		Val:  Book{},
	}
	return BookShelf{
		DummyHead: bn,
		Last:      bn,
	}
}

func (bs *BookShelf) AppendBook(book Book) {
	bookNode := InitBookNode(book)
	bs.Last.Next = bookNode
	bs.Last = bs.Last.Next
}

func (bs *BookShelf) Iterator() interface{} {
	return InitBookShelfIterator(bs)
}
