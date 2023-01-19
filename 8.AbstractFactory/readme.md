# AbstractFactory

**抽象工厂的工作是将”抽象零件“组装成”抽象产品“**

我们无需关心零件的具体实现，而是只关心接口。仅使用接口将零件组装为产品。

# 实例

仅包含关键代码

具体代码在Github上查看

## Factory接口

```golang
type Factory interface {
   CreateLink(string, string) Link
   CreateTray(string) Tray
   CreatePage(string, string) Page
}
```

## Item接口

```golang
type Item interface {
   MakeHtml() string
}
```

##  Factory方法

```golang
func GetFactory(way string) _interface.Factory {
   switch way {
   case "list":
      return list.Factory{}
   case "table":
      return table.Factory{}
   default:
      fmt.Println("暂无此工厂")
      return nil
   }
}
```

## ListFactory结构体

```golang
type Factory struct {
}

func (f Factory) CreateLink(url, caption string) abstract_parts.Link {
   return parts.InitLink(url, caption)
}
func (f Factory) CreateTray(caption string) abstract_parts.Tray {
   return parts.InitTray(caption)
}
func (f Factory) CreatePage(title, name string) abstract_parts.Page {
   return parts.InitPage(title, name)
}
```

## 拓展思路

## 易于增加具体的工厂

* AbstractFactory模式增加具体的工厂相当容易，指需要编写那些类实现哪些方法非常清楚
* 在实例中新增工厂只需要实现子类Factory、Link、Tray、Page，并定义方法，再注册组件即可。

