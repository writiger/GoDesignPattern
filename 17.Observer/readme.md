# Observer模式

​	在Observer模式中，当观察对象的状态发生变化时，会通知给观察者。Observer模式适用于根据对象状态进行相应处理的场景。

# 实例

## NumberGenerator接口

```golang
type NumberGenerator interface {
   AddObserver(Observer)
   DeleteObserver(Observer)
   NotifyObservers()
   GetNumber() int
   Execute()
}
```

## Observer接口

```golang
type Observer interface {
   Update(NumberGenerator)
}
```

## DigitObserver结构体

```golang
type DigitObserver struct {
}

func (do DigitObserver) Update(generator _interface.NumberGenerator) {
   fmt.Println("DigitObserver: ", generator.GetNumber())
   time.Sleep(time.Millisecond * 100)
}
```

## GraphObserver结构体

```golang
type GraphObserver struct {
}

func (_go GraphObserver) Update(generator _interface.NumberGenerator) {
   fmt.Print("GraphObserver: ")
   for i := 0; i < generator.GetNumber(); i++ {
      fmt.Print("*")
   }
   fmt.Println()
   time.Sleep(time.Millisecond * 100)
}
```

## RandomNumberGenerator结构体

```golang
type RandomNumberGenerator struct {
   number    int
   observers []_interface.Observer
}

func (rng *RandomNumberGenerator) AddObserver(observer _interface.Observer) {
   rng.observers = append(rng.observers, observer)
}

func (rng *RandomNumberGenerator) DeleteObserver(observer _interface.Observer) {
   flag := -1
   for i := 0; i < len(rng.observers); i++ {
      if observer == rng.observers[i] {
         flag = i
      }
   }
   if flag != -1 {
      rng.observers = append(rng.observers[:flag], rng.observers[flag+1:]...)
   }
}

func (rng *RandomNumberGenerator) NotifyObservers() {
   for _, item := range rng.observers {
      item.Update(rng)
   }
}

func (rng *RandomNumberGenerator) GetNumber() int {
   return rng.number
}

func (rng *RandomNumberGenerator) Execute() {
   rand.Seed(time.Now().UnixNano())
   for i := 0; i < 20; i++ {
      rng.number = rand.Intn(50)
      rng.NotifyObservers()
   }
}
```

# 拓展思路

## 可替换性

* 利用接口抽象出方法
* 使用接口作为参数、返回值以及保存实例

## 从”观察“变为”通知“

​	Observer本意为”观察者“，但实际上Observer角色并非主动去观察，而是被动接受来自Subject角色的通知。因此，Observer模式也被称为Pulish-Subscribe（发布-订阅）模式。