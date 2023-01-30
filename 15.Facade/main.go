package main

import (
	"15.Facade/facade"
	"fmt"
)

func main() {
	err := facade.MakeWelcomePage("zhangsan@qq.com", "welcome.html")
	if err != nil {
		fmt.Println(err.Error())
	}
}
