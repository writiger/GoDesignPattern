package main

import (
	"Interator/books"
	"fmt"
)

func main() {
	book1 := books.InitBook("图解设计模式")
	book2 := books.InitBook("代码整洁之道")
	book3 := books.InitBook("大话数据结构")
	bookShell := books.InitBookShell()
	bookShell.AppendBook(book1)
	bookShell.AppendBook(book2)
	bookShell.AppendBook(book3)
	it := books.InitBookShelfIterator(&bookShell)
	for it.HasNext() {
		book := it.Next().(books.Book)
		fmt.Println(book.GetName())
	}
}
