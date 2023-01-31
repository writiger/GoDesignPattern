package observer

import (
	_interface "17.Observer/interface"
	"math/rand"
	"time"
)

type RandomNumberGenerator struct {
	number    int
	observers []_interface.Observer
}

func (rng *RandomNumberGenerator) AddObserver(observer _interface.Observer) {
	rng.observers = append(rng.observers, observer)
}

func (rng *RandomNumberGenerator) DeleteObserver(observer _interface.Observer) {
	flag := -1
	for i := 0; i < len(rng.observers); i++ {
		if observer == rng.observers[i] {
			flag = i
		}
	}
	if flag != -1 {
		rng.observers = append(rng.observers[:flag], rng.observers[flag+1:]...)
	}
}

func (rng *RandomNumberGenerator) NotifyObservers() {
	for _, item := range rng.observers {
		item.Update(rng)
	}
}

func (rng *RandomNumberGenerator) GetNumber() int {
	return rng.number
}

func (rng *RandomNumberGenerator) Execute() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		rng.number = rand.Intn(50)
		rng.NotifyObservers()
	}
}
