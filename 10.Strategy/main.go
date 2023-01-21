package main

import (
	"10.Strategy/strategy"
	"fmt"
)

func main() {
	p1 := strategy.InitPlayer("Tom", strategy.InitWinning())
	p2 := strategy.InitPlayer("Jerry", strategy.InitProb())
	for i := 0; i < 10000; i++ {
		h1 := p1.NextHand()
		h2 := p2.NextHand()
		if h1.IsStrongerThan(h2) {
			p1.Win()
			p2.Lose()
			fmt.Println("Winner" + p1.ToString())
		} else if h1.IsWeakerThan(h2) {
			p2.Win()
			p1.Lose()
			fmt.Println("Winner" + p2.ToString())
		} else {
			p1.Even()
			p2.Even()
			fmt.Println("Even...")
		}
	}
	fmt.Println("Total Result:")
	fmt.Println(p1.ToString())
	fmt.Println(p2.ToString())
}

func play(p1, p2 *strategy.Player) string {
	h1 := p1.NextHand()
	h2 := p2.NextHand()
	if h1.IsStrongerThan(h2) {
		p1.Win()
		p2.Lose()
		return "Winner" + p1.ToString()
	} else if h1.IsWeakerThan(h2) {
		p2.Win()
		p1.Lose()
		return "Winner" + p2.ToString()
	} else {
		return "Even...\n"
	}
}
