package print

import "fmt"

type Banner struct {
	Info string
}

func (b *Banner) ShowWithParen() {
	fmt.Println("(", b.Info, ")")
}

func (b *Banner) ShowWithAster() {
	fmt.Println("*", b.Info, "*")
}
