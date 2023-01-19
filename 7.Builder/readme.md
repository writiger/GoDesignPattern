# Builder 模式

在建造大楼时，需要先打牢地基，搭建框架，然后自下而上地一层一层盖起来。通常，在建造这种具有复杂结构的物体时，很难一气呵成。我们首先需要建造组成这个物体的各个部分，然后分阶段将他们组装起来。

# 实例

我们将编写一段用于编写“文档”的实例程序。

这里的文档具有以下结构：

* 含有一个标题
* 含有几个字符串
* 含有条目项目

我们将实现两种Builder

*  TextBuilder
*  HtmlBuilder

## Director接口

```golang
type Director interface {
   Construct()
}
```

## Builder接口

```golang
type Builder interface {
   MakeTitle(string)
   MakeString(string)
   MakeItems(...string)
   Close()
}
```

## Director结构体

```golang
type Director struct {
   builder _interface.Builder
}

func InitDirector(b _interface.Builder) *Director {
   return &Director{builder: b}
}

func (d *Director) Construct() {
   d.builder.MakeTitle("Greeting")
   d.builder.MakeString(" 从早上到下午")
   d.builder.MakeItems("早上好", "下午好")
   d.builder.MakeString("晚上")
   d.builder.MakeItems("晚上好", "晚安", "再见")
   d.builder.Close()
}
```

## HtmlBuilder结构体

```golang
type HtmlBuilder struct {
   name  string
   fileP *os.File
}

func InitHtmlBuilder(name string) *HtmlBuilder {
   filename := name + ".html"
   var f *os.File
   if _, err := os.Stat(filename); err != nil {
      f, _ = os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
   }
   f, _ = os.Create(filename)
   return &HtmlBuilder{name: name, fileP: f}
}

func (h *HtmlBuilder) MakeTitle(s string) {
   h.fileP.WriteString("<html><head><title>" + h.name + "</title></head><body>")
   h.fileP.WriteString("<h1>" + h.name + "</h1>")
}

func (h *HtmlBuilder) MakeString(s string) {
   h.fileP.WriteString("<p>" + s + "</p>")
}

func (h *HtmlBuilder) MakeItems(items ...string) {
   h.fileP.WriteString("<ul>")
   for _, item := range items {
      h.fileP.WriteString("<li>" + item + "</li>")
   }
   h.fileP.WriteString("</ul>")
}

func (h *HtmlBuilder) Close() {
   h.fileP.WriteString("</body></html>")
   defer h.fileP.Close()
}

func (h *HtmlBuilder) GetResult() string {
   return h.name
}
```

## TextBuilder结构体

```golang
type TextBuilder struct {
   builder strings.Builder
}

func InitTextBuilder() *TextBuilder {
   return &TextBuilder{}
}

func (t *TextBuilder) MakeTitle(s string) {
   t.builder.WriteString("=============================\n")
   t.builder.WriteString("[" + s + "]\n")
   t.builder.WriteString("\n")
}

func (t *TextBuilder) MakeString(s string) {
   t.builder.WriteString("■" + s + "\n")
   t.builder.WriteString("\n")
}

func (t *TextBuilder) MakeItems(items ...string) {
   for _, item := range items {
      t.builder.WriteString("· " + item + "\n")
   }
   t.builder.WriteString("\n")
}

func (t *TextBuilder) Close() {
   t.builder.WriteString("=============================\n")
}

func (t *TextBuilder) GetResult() string {
   return t.builder.String()
}
```

# 拓展思路

## 谁知道什么

在面向对象编程中，“谁知道什么”是非常重要的。

Director类不知道自己使用的究竟是那个`Builder`。**只有不知道才能替换**

真是因为可替换性组件才具有高价值，作为设计人员，我们必须时刻关注这种“**可替换性**”

## 设计时能够决定的事情和不能决定的事情

在示例程序中我们只实现了纯文本类的Builder和Html类的Builder，但未来可能需要编写新的Builder，到时候需要设计新的方法吗？

显然类的设计者不是神仙，他们无法预测将来的事。但是，我们还是要尽可能灵活地应对近期可能发生的变化。

