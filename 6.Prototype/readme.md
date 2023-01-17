# Prototype 模式

开发过程中，有时会有“在不指定类名的前提下生成实例”的需求。

例如，以下情况，我们就不能根据类生成实例

1. **对象类型种类繁多，无法整合到一个类中**

   需要处理的对象太多，如果将他们分别作为一个类，必须需要编写多个类文件

2. **难以根据类生成实例**

   生成的实例太过复杂，很难根据类来生成实例。比如，假设这里有一个，用户通过图形化界面编辑器使用鼠标制作出的图形实例。想在程序中创建这样的实例是非常困难的。通常，想生成一个完全一样的实例时，会将用户通过操作创建的实例保存起来，在需要的时候复制生成新的实例。

3. **想解耦框架与生成实例时**

   想要生成实例的框架不依赖与具体的类。这时，不能指定类名来生成实例，要事先“注册”一个“原型”实例，然后通过复制实例来生成新的实例。

# 实例

## Product接口

```golang
type Product interface {
   Use(string)
   Cloneable
}

type Cloneable interface {
   Clone() Product
}
```

## Manager结构体

```golang
type Manager struct {
   Showcase map[string]_interface.Product
}

func (m *Manager) Register(name string, pro _interface.Product) {
   if _, ok := m.Showcase[name]; ok {
      return
   }
   m.Showcase[name] = pro
}

func (m *Manager) Create(name string) _interface.Product {
   if pro, ok := m.Showcase[name]; ok {
      return pro.Clone()
   }
   return nil
}
```

## MessageBox结构体

```golang
type MessageBox struct {
   decochar string
}

func InitMessageBox(s string) MessageBox {
   return MessageBox{decochar: s}
}

func (m MessageBox) Use(s string) {
   length := len(s)
   for i := 0; i < length+4; i++ {
      fmt.Print(m.decochar)
   }
   fmt.Println()
   fmt.Println(m.decochar, s, m.decochar)
   for i := 0; i < length+4; i++ {
      fmt.Print(m.decochar)
   }
   fmt.Println()
}

func (m MessageBox) Clone() _interface.Product {
   return MessageBox{decochar: m.decochar}
}
```

## UnderlinePen结构体

```golang
type UnderlinePen struct {
   decochar string
}

func InitUnderlinePen(s string) UnderlinePen {
   return UnderlinePen{decochar: s}
}

func (u UnderlinePen) Use(s string) {
   fmt.Println(`"`, s, `"`)
   length := len(s)
   for i := 0; i < length+4; i++ {
      fmt.Print(u.decochar)
   }
   fmt.Println()
}

func (u UnderlinePen) Clone() _interface.Product {
   return UnderlinePen{decochar: u.decochar}
}
```