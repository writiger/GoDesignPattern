package observer

import (
	_interface "17.Observer/interface"
	"fmt"
	"time"
)

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
