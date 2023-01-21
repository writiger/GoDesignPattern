package strategy

import (
	_interface "10.Strategy/interface"
	"crypto/rand"
	"fmt"
	"math/big"
)

type Prob struct {
	history      [][]int
	preHandValue int
	curHandValue int
}

func InitProb() *Prob {
	h := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
	return &Prob{
		history:      h,
		preHandValue: 0,
		curHandValue: 0,
	}
}

func (p *Prob) getSum(hv int) int64 {
	sum := 0
	for _, i := range p.history[hv] {
		sum += i
	}
	return int64(sum)
}

// NextHand 通过一个表字段(history)
// 计算连续两局的胜负分配权重
// 根据权重生成随机数
func (p *Prob) NextHand() _interface.Hand {
	n, _ := rand.Int(rand.Reader, big.NewInt(p.getSum(p.curHandValue)))
	bet := int(n.Int64())
	var handValue int
	if bet < p.history[p.curHandValue][0] {
		handValue = 0
	} else if bet < p.history[p.curHandValue][0]+p.history[p.curHandValue][1] {
		handValue = 1
	} else {
		handValue = 2
	}
	p.preHandValue = p.curHandValue
	p.curHandValue = handValue
	fmt.Printf("jer:%s\n", Name[handValue])
	return InitHand(handValue)
}

func (p *Prob) Study(win bool) {
	if win {
		p.history[p.preHandValue][p.curHandValue]++
	} else { // 通过其他加一表示 当前减一
		p.history[p.preHandValue][(p.curHandValue+1)%3]++
		p.history[p.preHandValue][(p.curHandValue+2)%3]++
	}
}
