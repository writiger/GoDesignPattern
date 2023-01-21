package strategy

import _interface "10.Strategy/interface"

const (
	HandValueGUU = iota // 石头
	HandValueCHO        // 布
	HandValuePAA        // 剪刀
)

var Name = make([]string, 3)
var Hands = make([]Hand, 3)

func init() {
	Name = []string{
		"石头",
		"剪刀",
		"布",
	}
	Hands = []Hand{
		{HandValueGUU},
		{HandValueCHO},
		{HandValuePAA},
	}
}

type Hand struct {
	handValue int
}

func InitHand(i int) Hand {
	return Hands[i]
}

// 平：0 输：-1 赢：1
func (h Hand) fight(in _interface.Hand) int {
	in = in.(Hand)
	if h == in {
		return 0
	} else if (h.handValue+1)%3 == in.GetHandValue() {
		return -1
	} else {
		return 1
	}
}

func (h Hand) GetHandValue() int {
	return h.handValue
}

func (h Hand) IsStrongerThan(in _interface.Hand) bool {
	return h.fight(in) == 1
}
func (h Hand) IsWeakerThan(in _interface.Hand) bool {
	return h.fight(in) == -1
}
