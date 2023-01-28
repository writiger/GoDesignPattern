# ChainOfResponsibility模式

​	我们先看看什么是**推卸责任**。假设我们要去公司领取资料。首先我们要向公司前台打听要去哪里领资料，他告诉我们应该去“营业窗口”。等我们到了”营业窗口“后，又被告知去“售后部门”。等我们好不容易到了“售后部门”时，又被告知应该去“资料中心”。像这样，在找到合适的办事人之前，我们被不断地踢给一个又一个人，这就是“推卸责任”。

​	“推卸责任”听起来有些贬义，但有时候确实需要“推卸责任”的情况。例如，当外部请求程序进行某个处理，但程序暂时无法直接决定由那个对象处理时，就需要责任链模式。这种情况下，我们可以考虑将**多个对象组成一条职责链，然后按照他们在职责链上的顺序一个一个的找出来由谁负责**。

# 实例

## Support接口

```golang
type Support interface {
   SetNext(Support) Support
   GetNext() Support
   GetName() string
   Resolve(Trouble) bool
   Support(Trouble)
}
```

## Trouble接口

```golang
type Trouble interface {
   GetNumber() int
   ToString() string
}
```

## Trouble结构体

```golang
type Trouble struct {
   number int
}

func InitTrouble(n int) Trouble {
   return Trouble{n}
}

func (t Trouble) GetNumber() int {
   return t.number
}

func (t Trouble) ToString() string {
   return "[Trouble " + strconv.Itoa(t.number) + " ]"
}
```

## Support结构体单例

```golang
type Support struct {
}

var SupportImpl *Support

func init() {
   SupportImpl = &Support{}
}

func (Support) CommonResolve(support _interface.Support, trouble _interface.Trouble) {
   if support.Resolve(trouble) {
      done(support.GetName(), trouble.ToString())
   } else if next := support.GetNext(); next != nil {
      next.Support(trouble)
   } else {
      fail(trouble.ToString())
   }
}

func done(s, t string) {
   fmt.Println(t, "is resolved by [", s, "]")
}

func fail(t string) {
   fmt.Println(t, "cannot be resolved.")
}
```

## NoSupport结构体

```golang
type NoSupport struct {
   name string
   next _interface.Support
}

func InitNoSupport(name string) *NoSupport {
   return &NoSupport{
      name: name,
      next: nil,
   }
}

func (ns *NoSupport) SetNext(next _interface.Support) _interface.Support {
   ns.next = next
   return next
}

func (ns *NoSupport) GetNext() _interface.Support {
   return ns.next
}

func (ns *NoSupport) GetName() string {
   return ns.name
}

func (ns *NoSupport) Resolve(_interface.Trouble) bool {
   return false
}

func (ns *NoSupport) Support(t _interface.Trouble) {
   SupportImpl.CommonResolve(ns, t)
}
```

## LimitSupport结构体

```golang
type LimitSupport struct {
   name  string
   next  _interface.Support
   limit int
}

func InitLimitSupport(name string, limit int) *LimitSupport {
   return &LimitSupport{
      name:  name,
      next:  nil,
      limit: limit,
   }
}

func (ls *LimitSupport) SetNext(next _interface.Support) _interface.Support {
   ls.next = next
   return next
}

func (ls *LimitSupport) GetNext() _interface.Support {
   return ls.next
}

func (ls *LimitSupport) GetName() string {
   return ls.name
}

func (ls *LimitSupport) Resolve(t _interface.Trouble) bool {
   if t.GetNumber() >= ls.limit {
      return false
   }
   return true
}

func (ls *LimitSupport) Support(t _interface.Trouble) {
   SupportImpl.CommonResolve(ls, t)
}
```

## OddSupport结构体

```golang
type OddSupport struct {
   name string
   next _interface.Support
}

func InitOddSupport(name string) *OddSupport {
   return &OddSupport{
      name: name,
      next: nil,
   }
}

func (os *OddSupport) SetNext(next _interface.Support) _interface.Support {
   os.next = next
   return next
}

func (os *OddSupport) GetNext() _interface.Support {
   return os.next
}

func (os *OddSupport) GetName() string {
   return os.name
}

func (os *OddSupport) Resolve(t _interface.Trouble) bool {
   if t.GetNumber()%2 == 0 {
      return false
   }
   return true
}

func (os *OddSupport) Support(t _interface.Trouble) {
   SupportImpl.CommonResolve(os, t)
}
```

# 拓展思路

## 弱化了发出请求的人和处理请求的人的关系

* 如果不使用这个模式就需要有某个伟大的角色知道“谁应该交给谁处理”

## 专注于自己的工作

* 每个角色都专注于自己的工作，无法处理时，干脆地对下一个处理者说”交给你了“。