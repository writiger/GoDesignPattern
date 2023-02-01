package main

import (
	"18.Memento/game"
	"fmt"
	"strconv"
	"time"
)

func main() {
	gamer := game.InitGamer(100)
	memento := gamer.CreateMemento()
	for i := 0; i < 100; i++ {
		fmt.Println("====", strconv.Itoa(i))
		fmt.Println("当前状态: ", gamer)
		gamer.Bet()
		fmt.Println("所持有金钱为", gamer.GetMoney(), "元")
		if gamer.GetMoney() > memento.GetMoney() {
			fmt.Println("所持有金钱增加了许多，因此保存当前游戏状态")
			memento = gamer.CreateMemento()
		} else if gamer.GetMoney() < memento.GetMoney()/2 {
			fmt.Println("所持有金钱减少了许多，因此将游戏恢复到之前的状态")
			gamer.RestoreMemento(memento)
		}
		time.Sleep(time.Millisecond * 10)
		fmt.Println()
	}
}
