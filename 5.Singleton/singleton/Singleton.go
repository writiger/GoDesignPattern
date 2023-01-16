package singleton

import (
	"fmt"
)

type singleton struct {
}

var s1 *singleton

// 使用golang的init函数生成一个实例
// 将生成的实例地址赋值给上文生成的变量
// 每次调用InitSingleton函数都返回这个公共变量
func init() {
	fmt.Println("生成一个实例")
	s1 = new(singleton)
}

func InitSingleton() *singleton {
	// 每次调用InitSingleton都返回由init函数初始化的实例
	// 确保了任何情况下调用InitSingleton都只生成一个实例
	return s1
}
