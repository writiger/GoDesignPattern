package strategy

import (
	_interface "10.Strategy/interface"
	"crypto/rand"
	"fmt"
	"math/big"
)

type Winning struct {
	won     bool
	preHand Hand
}

func InitWinning() *Winning {
	return &Winning{won: false}
}

// NextHand 一直出相同的，直到输为止
func (s *Winning) NextHand() _interface.Hand {
	if !s.won {
		n, _ := rand.Int(rand.Reader, big.NewInt(3))
		bet := int(n.Int64())
		s.preHand = InitHand(bet)
	}
	fmt.Printf("tom:%s\n", Name[s.preHand.handValue])
	return s.preHand
}

func (s *Winning) Study(win bool) {
	s.won = win
}
