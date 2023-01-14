package books

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
