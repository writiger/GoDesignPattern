# Factory Method 模式

在FactoryMethod模式里，父类决定实例的生产方式，但不决定所要生成的具体的类，具体的处理全部交给子类负责。这样可以将生成实例的框架和负责生成实例的类解耦。

# 实例

## Product接口

```golang
type Product interface {
   Use()
}
```

## Factory接口

```golang
type Factory interface {
   CreateProduct(string) factory.IDCard
   RegisterProduct(string)
   Create(string) factory.IDCard
}
```

## InitFactory方法

```golang
func InitFactory(factoryName string) interface{} {
   switch factoryName {
   case "IDCard":
      return IDCardFactoryTemplate{}
   default:
      return nil
   }
}
```

## IDCardFactory结构体

```golang
// IDCardFactory 子类
type IDCardFactory struct {
   IDCardFactoryTemplate
}

// IDCardFactoryTemplate 父类
type IDCardFactoryTemplate struct {
   Owners []string
}

func (idcft *IDCardFactoryTemplate) CreateProduct(name string) IDCard {
   fmt.Println("制作" + name + "的IDCard")
   return IDCard{Owner: name}
}

func (idcft *IDCardFactoryTemplate) RegisterProduct(name string) {
   idcft.Owners = append(idcft.Owners, name)
}

// Create 具体处理
func (idcf *IDCardFactory) Create(name string) IDCard {
   idcf.RegisterProduct(name)
   return idcf.CreateProduct(name)
}

func InitIDCardFactory() IDCardFactory {
   return IDCardFactory{InitFactory("IDCard").(IDCardFactoryTemplate)}
}
```

## IDCard结构体

```golang
type IDCard struct {
   Owner string
}

func (idc *IDCard) Use() {
   fmt.Println("使用" + idc.Owner + "的IDCard")
}
```

