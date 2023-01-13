# Adapter模式

程序中经常会存在现有程序无法**直接使用**，需要做**适当变化**后才能使用的情况。
这种填补“现有程序”和“所需程序”之间差异的设计模式就是Adapter模式。

Adapter模式也被称为Wrapper模式。
Wrapper有” 包装器 “的意思替我们将某样东西包装起来，使其能够用于其他用途的东西就被称为”**包装器**“或者”**适配器**“

# 实例（使用委托

* 委托是指将某个方法的实际处理交给其他实例

## Print接口

```golang
type Print interface {
   PrintWeak()
   PrintStrong()
}
```

## Banner结构体

```golang
type Banner struct {
   Info string
}

func (b *Banner) ShowWithParen() {
   fmt.Println("(", b.Info, ")")
}

func (b *Banner) ShowWithAster() {
   fmt.Println("*", b.Info, "*")
}
```

## PrintBanner结构体

```golang
type PrintBanner struct {
   Banner Banner
}

func InitPrintBanner(info string) PrintBanner {
   return PrintBanner{
      Banner: Banner{Info: info},
   }
}

func (pb *PrintBanner) PrintWeak() {
   pb.Banner.ShowWithParen()
}

func (pb *PrintBanner) PrintStrong() {
   pb.Banner.ShowWithAster()
}
```

# 拓展思路

## 什么时候使用Adapter模式

* 重复利用已有的代码便于创建方法群
* 调试Bug时可以明确的知道不在现有`类`中

## 新老版本更替

* 使用Adapter模式实现新旧版本兼容

