package books

type BookShelfIterator struct {
	BookShelf BookShelf
	Index     *BookNode
}

func InitBookShelfIterator(shell *BookShelf) BookShelfIterator {
	return BookShelfIterator{
		BookShelf: *shell,
		Index:     shell.DummyHead.Next,
	}
}

func (bsi *BookShelfIterator) HasNext() bool {
	// HasNext方法可以理解为 “当前节点是否可以调用Next()”
	// 当前节点不为空时可以调用Next()返回true
	return bsi.Index != nil
	// 当前节点为nil时无法调用Next则返回false
}

func (bsi *BookShelfIterator) Next() interface{} {
	book := bsi.Index.Val
	// 向后移一位
	bsi.Index = bsi.Index.Next
	return book
}
