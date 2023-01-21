# Strategy模式

无论什么程序，其目的都是解决问题。而为了解决问题，我们又需要编写特定算法。使用Strategy模式可以整体地替换算法的实现部分。能够整体地替换算法，能够让我们轻松地以不同方法解决同一个问题，这种模式就是Strategy模式。

# 实例

## Strategy接口

```golang
type Strategy interface {
   NextHand() Hand
   Study(bool)
}
```

## Hand接口

```golang
type Hand interface {
   IsStrongerThan(Hand) bool
   IsWeakerThan(Hand) bool
   GetHandValue() int
}
```

## Hand.go

```golang
const (
   HandValueGUU = iota // 石头
   HandValueCHO        // 布
   HandValuePAA        // 剪刀
)

var Name = make([]string, 3)
var Hands = make([]Hand, 3)

func init() {
   Name = []string{
      "石头",
      "剪刀",
      "布",
   }
   Hands = []Hand{
      {HandValueGUU},
      {HandValueCHO},
      {HandValuePAA},
   }
}

type Hand struct {
   handValue int
}

func InitHand(i int) Hand {
   return Hands[i]
}

// 平：0 输：-1 赢：1
func (h Hand) fight(in _interface.Hand) int {
   in = in.(Hand)
   if h == in {
      return 0
   } else if (h.handValue+1)%3 == in.GetHandValue() {
      return -1
   } else {
      return 1
   }
}

func (h Hand) GetHandValue() int {
   return h.handValue
}

func (h Hand) IsStrongerThan(in _interface.Hand) bool {
   return h.fight(in) == 1
}
func (h Hand) IsWeakerThan(in _interface.Hand) bool {
   return h.fight(in) == -1
}
```

## Winning结构体

```golang
type Winning struct {
   won     bool
   preHand Hand
}

func InitWinning() *Winning {
   return &Winning{won: false}
}

// NextHand 一直出相同的，直到输为止
func (s *Winning) NextHand() _interface.Hand {
   if !s.won {
      n, _ := rand.Int(rand.Reader, big.NewInt(3))
      bet := int(n.Int64())
      s.preHand = InitHand(bet)
   }
   fmt.Printf("tom:%s\n", Name[s.preHand.handValue])
   return s.preHand
}

func (s *Winning) Study(win bool) {
   s.won = win
}
```

WinningStrategy是一个有些笨的策略——“上局赢了就继续出，否则随机出”

## Prob结构体

```golang
type Prob struct {
   history      [][]int
   preHandValue int
   curHandValue int
}

func InitProb() *Prob {
   h := [][]int{
      {1, 1, 1},
      {1, 1, 1},
      {1, 1, 1},
   }
   return &Prob{
      history:      h,
      preHandValue: 0,
      curHandValue: 0,
   }
}

func (p *Prob) getSum(hv int) int64 {
   sum := 0
   for _, i := range p.history[hv] {
      sum += i
   }
   return int64(sum)
}

// NextHand 通过一个表字段(history)
// 计算连续两局的胜负分配权重
// 根据权重生成随机数
func (p *Prob) NextHand() _interface.Hand {
   n, _ := rand.Int(rand.Reader, big.NewInt(p.getSum(p.curHandValue)))
   bet := int(n.Int64())
   var handValue int
   if bet < p.history[p.curHandValue][0] {
      handValue = 0
   } else if bet < p.history[p.curHandValue][0]+p.history[p.curHandValue][1] {
      handValue = 1
   } else {
      handValue = 2
   }
   p.preHandValue = p.curHandValue
   p.curHandValue = handValue
   fmt.Printf("jer:%s\n", Name[handValue])
   return InitHand(handValue)
}

func (p *Prob) Study(win bool) {
   if win {
      p.history[p.preHandValue][p.curHandValue]++
   } else { // 通过其他加一表示 当前减一
      p.history[p.preHandValue][(p.curHandValue+1)%3]++
      p.history[p.preHandValue][(p.curHandValue+2)%3]++
   }
}
```

ProbStrategy是一个具体策略，会根据以前的猜拳结果进行计算。

history字段是一个表，用于根据过去的胜负计算概率。它是一个二维数组，每个数组下标意思如下：

`history[上一局出的手势][这一局出的手势]`

**这个表达式的值越大，表示过去的胜率越高**

假设我们上一局出的石头

* `history[0][0]` 两句分别出石头、石头时的胜场
* `history[0][1]` 两句分别出石头、剪刀时的胜场
* `history[0][2]` 两句分别出石头、布时的胜场

那么我们就可以根据这三个表达式计算本局应该出什么的权重

例如，如果：

* `history[0][0]=3`
* `history[0][1]=5`
* `history[0][2]=7`

那么下一局石头、剪刀、布的权重分别为3、5、7和为15 。在[0,15) 中取一个随机数x

* 如果x ∈ [0,3) 则出石头
* 如果x ∈ [3,8) 则出剪刀
* 如果x ∈ [8,15) 则出布

## Player结构体

```golang
type Player struct {
   name     string
   winCnt   int
   loseCnt  int
   gameCnt  int
   strategy _interface.Strategy
}

func InitPlayer(name string, strategy _interface.Strategy) *Player {
   return &Player{
      name:     name,
      winCnt:   0,
      loseCnt:  0,
      gameCnt:  0,
      strategy: strategy,
   }
}

func (p *Player) NextHand() Hand {
   return p.strategy.NextHand().(Hand)
}

func (p *Player) Win() {
   p.strategy.Study(true)
   p.winCnt++
   p.gameCnt++
}
func (p *Player) Lose() {
   p.strategy.Study(false)
   p.loseCnt++
   p.gameCnt++
}
func (p *Player) Even() {
   p.gameCnt++
}

func (p *Player) ToString() string {
   return fmt.Sprintf("[ %s : %d games, %d win, %d lose]\n", p.name, p.gameCnt, p.winCnt, p.loseCnt)
}
```

## 拓展思路

## 为什么需要特意编写Strategy接口

* 使用这种方法可以方便的替换算法