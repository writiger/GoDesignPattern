# Bridge模式

“Birdge”的意思是桥梁。就像现实世界中，桥梁的意思是将河流两侧连接起来一样，Bridge模式的作用也是将两种东西连接起来:

* 类的功能层次结构
* 类的实现结构

## 类的层次结构的两种作用

### 希望新增功能时

假设现在有一个类Something。当我们想在Something中增加新功能时（想增加一个具体方法时），会编写一个Something的子类（派生类），即SomethingGood类。这样就构成了一个小小的类层次结构。

* 父类具有基本功能
* 子类中新增功能

以上这种层次结构被称为”**类的功能层次结构**“

***注意***： **通常来说，类的层次结构关系不应过深**

### 希望增加新的实现时

父类声明接口，而子类的任务是实现接口。正是由于父类和子类的这种任务分担，我们才可以编写出高可替换性的类。

当子类ConcreteClass实现了父类AbstractClass类的接口时，他们之间就构成了一个小小的层次结构。

这种层次结构并非用于方便我们新增方法，他真正的作用时帮助我们实现下面的任务分组：

* 父类声明接口
* 子类实现接口

这种层次结构被称为**”类的实现层次结构“**



### 类的层次结构的混杂与分离

​	当我们编写子类时，需要确认自己的意图：”我是要增加功能呢？还是增加实现呢？“，当类的层次结构只有一层时，功能层次与实现层次是混杂在一个层次结构中的。这样很容易使类的结构层次变复杂，也难以透彻理解。

​	因此，我们将使用Bridge模式**分离两种层次结构**	

# 实例

## Display接口

```golang
type DisplayImpl interface {
   RawOpen()
   RawPrint()
   RawClose()
}
```

## Display 结构体

```golang
type Display struct {
   impl _interface.DisplayImpl
}

func InitDisplay(d _interface.DisplayImpl) *Display {
   return &Display{d}
}

func (d Display) open() {
   d.impl.RawOpen()
}

func (d Display) print() {
   d.impl.RawPrint()
}

func (d Display) close() {
   d.impl.RawClose()
}
```

## CountDisplay 结构体

```golang
// CountDisplay 新增功能
type CountDisplay struct {
   Display
}

func InitCountDisplay(impl _interface.DisplayImpl) *CountDisplay {
   return &CountDisplay{Display: *InitDisplay(impl)}
}

func (cd CountDisplay) MultiDisplay(times int) {
   cd.open()
   for i := 0; i < times; i++ {
      cd.print()
   }
   cd.close()
}
```

## StringDisplay

```golang
// StringDisplay 新增实现
type StringDisplay struct {
   info  string
   width int
}

func InitStringDisplay(info string) *StringDisplay {
   return &StringDisplay{
      info:  info,
      width: len(info),
   }
}

func (sd StringDisplay) RawOpen() {
   sd.printLine()
}

func (sd StringDisplay) RawPrint() {
   fmt.Println("|" + sd.info + "|")
}

func (sd StringDisplay) RawClose() {
   sd.printLine()
}

func (sd StringDisplay) printLine() {
   fmt.Print("+")
   for i := 0; i < sd.width; i++ {
      fmt.Print("-")
   }
   fmt.Println("+")
}
```

# 拓展思路

## 分开后更容易拓展

* 想要增加功能时，只需要对功能层次进行修改。”**增加后的功能可以被所有实现使用**“

* 例如：我们可以将类的”功能层级结构“应用于软件所运行的操作系统上，我们只需要编写一个操作系统共同的API角色，然后编写各个版本的实现。这样只需在功能层级新增功能，就可以在多个操作系统中实现。