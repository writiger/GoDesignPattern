# Visitor模式

​	在Visitor模式中，**数据结构和处理**被分离开来。我们需要编写一个“访问者”类来访问数据结构中的元素，并把对个元素处理交给访问者类。这样，当需要增加新的处理时，我们只需要编写新的访问者，然后让数据结构可以接受访问者的访问即可。

# 实例

## Element接口

```golang
type Element interface {
   Accept(Visitor)
}
```

## Entry接口

```golang
type Entry interface {
   GetName() string
   GetSize() int
   Element
}
```

## File接口

```golang
type File interface {
	Entry
}
```

## Directory接口

```golang
type Directory interface {
   Entry
   Add(Entry)
   GetRange() []Entry
}
```

## Visitor接口

```golang
type Visitor interface {
   VisitFile(File)
   VisitDirectory(Directory)
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

func (f File) GetName() string {
   return f.name
}

func (f File) GetSize() int {
   return f.size
}

func (f File) Accept(visitor _interface.Visitor) {
   visitor.VisitFile(f)
}
```

## Directory结构体

```golang
type Directory struct {
   name string
   dir  []_interface.Entry
}

func InitDirectory(name string) *Directory {
   return &Directory{
      name: name,
      dir:  make([]_interface.Entry, 0),
   }
}

func (d *Directory) GetRange() []_interface.Entry {
   return d.dir
}

func (d *Directory) GetName() string {
   return d.name
}

func (d *Directory) GetSize() int {
   res := 0
   for _, item := range d.dir {
      res += item.GetSize()
   }
   return res
}

func (d *Directory) Add(entry _interface.Entry) {
   d.dir = append(d.dir, entry)
}

func (d *Directory) Accept(visitor _interface.Visitor) {
   visitor.VisitDirectory(d)
}
```

## ListVisitor结构体

```golang
type Visitor struct {
   currentDir string
}

func InitVisitor() *Visitor {
   return &Visitor{currentDir: ""}
}

func (v *Visitor) VisitFile(f _interface.File) {
   fmt.Println(v.currentDir + "/" + entryToString(f))
}
func (v *Visitor) VisitDirectory(d _interface.Directory) {
   fmt.Println(v.currentDir + "/" + entryToString(d))
   saveDir := v.currentDir
   v.currentDir = v.currentDir + "/" + d.GetName()
   for _, item := range d.GetRange() {
      item.Accept(v)
   }
   v.currentDir = saveDir
}

func entryToString(e _interface.Entry) string {
   return e.GetName() + " (" + strconv.Itoa(e.GetSize()) + ")"
}
```

# 拓展思路

## 双重分发

​	我们来整理一下Visitor模式中的调用关系：

* accept方法调用如下：

  `element.accept(vistor)`

* visitor方法调用如下：

  `visitor.visit(element)`

对比一下会发现，这两个方法是相反的关系。

​	在Visitor模式中，ConcreteElement和ConcreteVisitor这两个角色共同决定了实际进行的处理。这种消息分发的方法一般被称为**双重分发**。

## 为什么要弄得这么复杂

​	当看到上面的处理流程后，大家可能会感觉到“Visitor模式不是把简单问题复杂化了么？”“如果处理循环，为什么不直接在结构体中编写循环语句呢？”

​	Visitor模式的目的是将**处理从数据结构中分离出来**。数据结构很重要，它能将元素集合关联在一起。但是需要注意的是，保存数据结构与以数据结构为基础进行处理是两种不同的东西。

​	Visitor模式提高了File类和Directory类**作为组件的独立性**。当每次要拓展功能，增加新的“处理”时，就不需要去修改File类和Directory类。

## 开闭原则——对拓展开放，对修改关闭

* **在不修改现有代码的前提下进行拓展**