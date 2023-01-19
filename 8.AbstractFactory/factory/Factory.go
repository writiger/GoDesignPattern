package factory

import (
	"8.AbstractFactory/factory/list"
	_interface "8.AbstractFactory/interface/abstract_parts"
	"fmt"
)

func GetFactory(way string) _interface.Factory {
	switch way {
	case "list":
		return list.Factory{}
	case "table":
		return nil
	default:
		fmt.Println("暂无此工厂")
		return nil
	}
}
