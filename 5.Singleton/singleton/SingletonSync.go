package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once

var s *singleton

// InitSingletonOnce 不会再项目初始化时创建实例
// 使用Once确保只创建一个实例
func InitSingletonOnce() *singleton {
	once.Do(func() {
		fmt.Println("创建一个实例")
		s = &singleton{}
	})
	return s
}
