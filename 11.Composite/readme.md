# Composite模式

​	在计算机的文件系统中，有“文件夹”的概念。文件夹里面既可以放入文件，也可以放入其他文件夹（子文件夹）。在子文件夹中，一样地即可以放入文件，又可以放入文件夹。可以说，文件夹是形成了一种容器结构、递归结构。

​	虽然文件夹与文件是不同类型的对象，但他们都“可以被放入文件夹中”。文件夹和文件有时也被统称为“目录条目”。在目录条目中，文件夹和文件被当作同一种对象看待。

​	将容器和内容作为同一种东西看待，帮助我们方便处理问题。在容器中即可以放入内容，也可以放入小容器，小容器中可以放入更小的容器。这样，就形成了容器结构、递归结构。**能够使容器和内容具有一致性，创造出递归结构**的模式就是Composite模式。

# 实例

## Entry接口

```golang
type Entry interface {
   GetName() string
   GetSize() int
   PrintList(string)
}
```

## Entry.go

```golang
func EntryToString(e _interface.Entry) string {
   return e.GetName() + " (" + strconv.Itoa(e.GetSize()) + ")"
}
```

## File结构体

```golang
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
```

## Directory结构体

```golang
type Directory struct {
   name      string
   directory []*_interface.Entry
}

func InitDirectory(name string) *Directory {
   return &Directory{
      name:      name,
      directory: make([]*_interface.Entry, 0),
   }
}

func (d *Directory) Add(entry _interface.Entry) {
   d.directory = append(d.directory, &entry)
}

func (d *Directory) GetName() string {
   return d.name
}

func (d *Directory) GetSize() int {
   res := 0
   for _, item := range d.directory {
      res += (*item).GetSize()
   }
   return res
}
```

# 拓展思路

## 多个和单个的一致性

​	使用Composite模式可以使容器与内容具有一致性，也可称为**多个和单个的一致性**

例如，以下为测试程序行为的场景：

* Test1用来测试键盘输入数据
* Test2用来测试文件输入数据
* Test3用来测试网络输入数据

我们可以将这些测试统一成“输入测试”