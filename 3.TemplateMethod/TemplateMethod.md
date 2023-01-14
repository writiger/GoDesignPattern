# Template Method 模式

带有模板功能的模式

`在父类中定义好处理流程的框架，在子类中实现`

# 实例

## AbstractDisplay接口

```golang
type AbstractDisplay interface {
   Open()
   Print()
   Close()
}

type Display interface {
   Display()
}
```

## Display方法

```golang
// Display 模板
func Display(d _interface.AbstractDisplay) {
   d.Open()
   for i := 0; i < 5; i++ {
      d.Print()
   }
   d.Close()
}
```

## CharDisplay结构体

```golang
type CharDisplay struct {
   Info string
}

type charDisplayTemplate struct {
   Info string
}

func (d charDisplayTemplate) Open() {
   fmt.Print("<<")
}
func (d charDisplayTemplate) Print() {
   fmt.Print(d.Info)
}
func (d charDisplayTemplate) Close() {
   fmt.Println(">>")
}

func (cd *CharDisplay) Display() {
   Display(charDisplayTemplate{Info: cd.Info})
}

func InitCharDisplay(info string) CharDisplay {
   return CharDisplay{Info: info}
}
```

## StringDisplay结构体

```golang
type StringDisplay struct {
   Info string
}

type stringDisplayTemplate struct {
   Info string
   Len  int
}

func (d stringDisplayTemplate) Open() {
   d.Println()
}
func (d stringDisplayTemplate) Print() {
   fmt.Println("|" + d.Info + "|")
}
func (d stringDisplayTemplate) Close() {
   d.Println()
}
func (d stringDisplayTemplate) Println() {
   fmt.Print("+")
   for i := 0; i < d.Len; i++ {
      fmt.Print("-")
   }
   fmt.Println("+")
}

func (sd *StringDisplay) Display() {
   Display(stringDisplayTemplate{
      Info: sd.Info,
      Len:  len(sd.Info),
   })
}

func InitStringDisplay(info string) StringDisplay {
   return StringDisplay{Info: info}
}
```

# 拓展思路

## 可以使逻辑处理通用化

* 无需在每个子类中编写算法

# 父类子类协作

* 在看不到模板的情况下很难编写子类