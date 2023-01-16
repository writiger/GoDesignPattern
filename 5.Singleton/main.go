package main

import (
	"5.Singleton/singleton"
	"fmt"
)

func main() {
	fmt.Println("start")
	// s1 s2已经是指针了 指向同一个singleton
	s1 := singleton.InitSingleton()
	s2 := singleton.InitSingleton()
	fmt.Println(s1 == s2)
	s3 := singleton.InitSingletonOnce()
	s4 := singleton.InitSingletonOnce()
	fmt.Println(s3 == s4)
}
