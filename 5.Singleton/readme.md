# Singleton模式

程序运行时通常会创建很多实例，但是，当我们想在程序中表示某个东西只存在一个时，就会存在”只能创建一个实例”的需求。



* 确保任何情况下只有一个实例
* 程序上表现出“只存在一个实例”

# 实例

## Singleton结构体|饿汉模式

~~~ golang
type singleton struct{}

var s *singleton

// 使用golang的init函数生成一个实例
// 将生成的实例地址赋值给上文生成的变量
// 每次调用InitSingleton函数都返回这个公共变量
func init() {
	fmt.Println("生成一个实例")
	s = new(singleton)
}

func InitSingleton() *singleton {
	// 每次调用InitSingleton都返回由init函数初始化的实例
	// 确保了任何情况下调用InitSingleton都只生成一个实例
	return s
}
~~~

## SingletonSync结构体|懒汉模式

```golang
var once sync.Once

var s *singleton

// InitSingletonOnce 不会再项目初始化时创建实例
// 使用Once确保只创建一个实例
func InitSingletonOnce() *singleton {
   once.Do(func() {
      fmt.Println("创建一个实例")
      s = &Singleton{}
   })
   return s
}
```

# 拓展思路

## 为什么必须设置限制

* 当存在多个实例时，实例互相影响可能产生意想不到的bug

## 注意

* 使用单例模式的结构体需要小写开头
* 否则其他包能够使用结构体名直接生成实例
* 单例模式将会失去意义