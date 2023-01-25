# Decorator模式

​	假如现在有一块蛋糕，如果只涂上奶油，其他什么都不做，就是奶油蛋糕。如果加上草莓，就是草莓奶油蛋糕。如果再加上一块黑色巧克力板，上面用白色巧克力写上姓名，然后插上代表年龄的蜡烛，就变成了一块生日蛋糕。

​	不论是蛋糕、奶油蛋糕、草莓蛋糕还是生日蛋糕，他们的核心都是蛋糕。不过，经过涂上奶油、加上草莓等装饰后，蛋糕的味道变得更加甜美了，目的更加明确了。

​	程序中的对象与蛋糕十分相似。首先有一个相当于蛋糕的对象，然后像不断地装饰蛋糕一样地添加对应功能，它就变成使用目的更加明确的对象。

​	像这样不断地为对象添加装饰的设计模式被称为**Decorator模式**。Decorator指的是“装饰物”。

# 实例

## Display接口

```golang
type Display interface {
   GetColumns() int
   GetRows() int
   GetRowText(int) string
}

type Show interface {
   Show(Display)
}
```

## StringDisplay结构体

```golang
type StringDisplay struct {
   info string
}

func InitStringDisplay(info string) *StringDisplay {
   return &StringDisplay{info: info}
}

func (sd StringDisplay) GetColumns() int {
   return len(sd.info)
}

func (sd StringDisplay) GetRows() int {
   return 1
}

func (sd StringDisplay) GetRowText(row int) string {
   if row != 0 {
      return ""
   }
   return sd.info
}
```

## SideBorder结构体

```golang
type SideBorder struct {
   display _interface.Display
   char    string
}

func InitSideBorder(display _interface.Display, char string) *SideBorder {
   return &SideBorder{
      display: display,
      char:    char,
   }
}

func (sb SideBorder) GetColumns() int {
   return 1 + sb.display.GetColumns() + 1
}

func (sb SideBorder) GetRows() int {
   return sb.display.GetRows()
}

func (sb SideBorder) GetRowText(row int) string {
   return sb.char + sb.display.GetRowText(row) + sb.char
}
```

## FullBorder结构体

```golang
type FullBorder struct {
   display _interface.Display
}

func InitFullBorder(display _interface.Display) *FullBorder {
   return &FullBorder{display: display}
}

func (fb FullBorder) GetColumns() int {
   return 1 + fb.display.GetColumns() + 1
}

func (fb FullBorder) GetRows() int {
   return 1 + fb.display.GetRows() + 1
}

func (fb FullBorder) GetRowText(row int) string {
   res := strings.Builder{}
   if row == 0 || row == fb.display.GetRows()+1 {
      res.WriteString("+")
      res.WriteString(makeLine("-", fb.display.GetColumns()))
      res.WriteString("+")
   } else {
      res.WriteString("|")
      res.WriteString(fb.display.GetRowText(row - 1))
      res.WriteString("|")
   }
   return res.String()
}

func makeLine(char string, count int) string {
   buf := strings.Builder{}
   for i := 0; i < count; i++ {
      buf.WriteString(char)
   }
   return buf.String()
}
```

## Table结构体

```golang
type Table struct {
}

func (t Table) Show(display _interface.Display) {
   rows := display.GetRows()
   for i := 0; i < rows; i++ {
      fmt.Println(display.GetRowText(i))
   }
}
```

# 拓展思路

## 在不改变装饰物的前提下增加功能

* 装饰物和被装饰物具有相同接口
* Decorator使用了委托，例如，SideBorder的`GetColumns`方法就会调用`display.GetColumns()`

## 可以动态的增加功能

不用改变框架代码，就可以生成一个与其他对象具有不同关系的新对象。

## 导致增加很多很小的类

* Decorator的一个缺点就是会导致程序中增加许多功能类似的很小的类