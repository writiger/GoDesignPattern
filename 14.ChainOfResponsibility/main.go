package main

import "14.ChainOfResponsibility/responsibility"

func main() {
	alice := responsibility.InitNoSupport("Alice")
	bob := responsibility.InitLimitSupport("Bob", 100)
	diana := responsibility.InitLimitSupport("Diana", 200)
	elmo := responsibility.InitOddSupport("Elmo")
	alice.SetNext(bob).SetNext(diana).SetNext(elmo)
	for i := 0; i < 500; i += 33 {
		alice.Support(responsibility.InitTrouble(i))
	}
}
