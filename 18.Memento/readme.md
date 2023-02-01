# Memento模式

​	我们在使用文本编辑器编写文件时，如果不小心删除了某句话，可以通过撤销功能将文件恢复至之前的状态。有些文本编辑器甚至支持多次撤销，能够恢复至很久之前的版本。

​	使用面向对象的编程方式实现撤销功能时，需要实现保存实例的相关状态信息。然后，在撤销时，还需要根据所保存的信息将实例恢复至原来的状态。

​	使用Memento模式能够实现应用程序的以下功能。

* Undo（撤销）
* Redo（重做）
* History（历史记录）
* Snapshot（快照）

# 实例

## Gamer结构体

```golang
type Gamer struct {
   money  int
   fruits []string
}

func InitGamer(money int) *Gamer {
   return &Gamer{
      money:  money,
      fruits: make([]string, 0),
   }
}

func (g *Gamer) GetMoney() int {
   return g.money
}

func (g *Gamer) Bet() {
   rand.Seed(time.Now().UnixNano())
   dice := rand.Intn(6) + 1
   if dice == 1 {
      g.money += 100
      fmt.Println("所持有金钱增加了。")
   } else if dice == 2 {
      g.money /= 2
      fmt.Println("所持有金钱减半了")
   } else if dice == 6 {
      fruit := getFruit()
      g.fruits = append(g.fruits, fruit)
      fmt.Println("获得了水果 （ ", fruit, " ）")
   } else {
      fmt.Println("什么都没有发生。")
   }
}

func (g *Gamer) CreateMemento() Memento {
   m := InitMemento(g.money)
   for _, item := range g.fruits {
      if strings.HasPrefix(item, "好吃的") {
         m.AddFruit(item)
      }
   }
   return m
}

func (g *Gamer) RestoreMemento(memento Memento) {
   g.money = memento.money
   g.fruits = memento.fruits
}

func (g *Gamer) String() string {
   return fmt.Sprintf("[money = %d , fruits = %v ]", g.money, g.fruits)
}

func getFruit() string {
   fruits := []string{
      "苹果",
      "葡萄",
      "香蕉",
      "橘子",
   }
   prefix := ""
   rand.Seed(time.Now().UnixNano())
   boolean := rand.Intn(1)
   fruit := rand.Intn(4)
   if boolean == 0 {
      prefix = "好吃的"
   }
   return prefix + fruits[fruit]
}
```

## Memento结构体

```golang
type Memento struct {
   money  int
   fruits []string
}

func InitMemento(money int) Memento {
   return Memento{
      money:  money,
      fruits: make([]string, 0),
   }
}

func (m *Memento) AddFruit(fruit string) {
   m.fruits = append(m.fruits, fruit)
}

func (m *Memento) GetFruit() []string {
   return m.fruits
}

func (m *Memento) GetMoney() int {
   return m.money
}
```

# 拓展思路

## 需要多少个Momento

​	在示例程序中，Main类只保存了一个Memento。如果使用切片等方式保存多个Memento类即可保存多个时间点的状态。

