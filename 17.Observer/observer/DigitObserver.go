package observer

import (
	_interface "17.Observer/interface"
	"fmt"
	"time"
)

type DigitObserver struct {
}

func (do DigitObserver) Update(generator _interface.NumberGenerator) {
	fmt.Println("DigitObserver: ", generator.GetNumber())
	time.Sleep(time.Millisecond * 100)
}
