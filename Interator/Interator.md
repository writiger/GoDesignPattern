# Iterator模式（迭代器

在golang中我们可以使用for循环遍历数组元素

~~~ golang
nums := []int{1, 2, 3, 4, 5}
for i:= 0;i< len(nums);i ++ {
    fmt.Println(nums[i])
}
~~~

* 像这样让`i`在循环中不断自增从而实现遍历的功能

* 这里将循环变量`i`的作用抽象化、通用化形成的模式在设计模式种被称为**Iterator模式**
* 下面是go中最基础的迭代器模式

~~~ golang
nums := []int{1, 2, 3, 4, 5}
for i, item := range nums {
    fmt.Printf("i : %d , item : %v\n", i, item)
}
~~~

# Go语言大多数程序不需要迭代器

因为内置的slice和map两种容器都可以通过range进行遍历，并且这两种容器在性能方面做了足够的优化。只要没有特殊的需求，通常是直接用这两种容器解决问题。即使不得不写了一个自定义容器，我们几乎总是可以实现一个函数，把所有元素（的引用）拷贝到一个slice之后返回，这样调用者又可以直接用range进行遍历了。

当然某些特殊场合迭代器还是有用武之地。比如迭代器的Next()是个耗时操作，不能一口气拷贝所有元素；再比如某些条件下需要中断遍历。

# 实例

* 这里使用经典方法的链表实现

* Go语言支持first class functions、高阶函数、闭包、多返回值函数。用上这些特性可以换种方式实现迭代器。



## Aggregate接口

```golang
type Aggregate interface {
   Iterator() interface{} // 返回一个用于遍历的迭代器
}
```

## Iterator接口

```golang
type Iterator interface {
   HasNext() bool // 判断是否存在下一个元素
   Next() interface{} // 返回当前元素并向下移动
}
```

## Book结构体

```golang
type Book struct {
   name string
}

// InitBook 模拟构造函数
// 结构体不需要放进堆中返回结构体优势更大
// 参考 https://mp.weixin.qq.com/s/PXGCqxK97U8mLGxW07ZTqw
func InitBook(nameInit string) Book {
   return Book{
      name: nameInit,
   }
}

// GetName 希望书名不允许被直接更改
func (b *Book) GetName() string {
   return b.name
}
```

## BookNode

```golang
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
```

## BookShelf结构体

```golang
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

```

## BookShelfIterator结构体

```golang
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

```

# 拓展思路

# 不管实现如何变化，都可以使用迭代器

* 引入Iterator后可以将遍历和实现分离开来
* 不管BookShelf的编写者采用什么管理书本BookShelf只要能够返回Iterator实例就可以实现遍历

## 难以理解接口

* 如果只是用具体的类来解决问题，很容易导致类之间的强耦合
* 要优先使用接口编程

## 迭代器的种类多种多样

* 从后向前
* 跳跃式遍历
* ...

