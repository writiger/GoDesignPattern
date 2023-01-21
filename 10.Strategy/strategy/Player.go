package strategy

import (
	_interface "10.Strategy/interface"
	"fmt"
)

type Player struct {
	name     string
	winCnt   int
	loseCnt  int
	gameCnt  int
	strategy _interface.Strategy
}

func InitPlayer(name string, strategy _interface.Strategy) *Player {
	return &Player{
		name:     name,
		winCnt:   0,
		loseCnt:  0,
		gameCnt:  0,
		strategy: strategy,
	}
}

func (p *Player) NextHand() Hand {
	return p.strategy.NextHand().(Hand)
}

func (p *Player) Win() {
	p.strategy.Study(true)
	p.winCnt++
	p.gameCnt++
}
func (p *Player) Lose() {
	p.strategy.Study(false)
	p.loseCnt++
	p.gameCnt++
}
func (p *Player) Even() {
	p.gameCnt++
}

func (p *Player) ToString() string {
	return fmt.Sprintf("[ %s : %d games, %d win, %d lose]\n", p.name, p.gameCnt, p.winCnt, p.loseCnt)
}
