package main

import "17.Observer/observer"

func main() {
	generator := observer.RandomNumberGenerator{}
	obs1 := observer.DigitObserver{}
	obs2 := observer.GraphObserver{}
	generator.AddObserver(obs1)
	generator.AddObserver(obs2)
	generator.Execute()
}
