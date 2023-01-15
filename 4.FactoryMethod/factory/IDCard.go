package factory

import "fmt"

type IDCard struct {
	Owner string
}

func (idc *IDCard) Use() {
	fmt.Println("使用" + idc.Owner + "的IDCard")
}
