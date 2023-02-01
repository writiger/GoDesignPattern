package game

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

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
