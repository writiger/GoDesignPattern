# Mediator模式

​	Mediator的意思是”仲裁者“”中介者“。一方面发生麻烦事的时候，会通知仲裁者；当发生涉及全部组员的事情时，也通知仲裁者。当仲裁者下达指示时，组员会立即执行。团队组员之间不再互相沟通私自做出决定，而是发生任何事情都向仲裁者报告。另一方面，仲裁者站在整个团队的角度上对组员上报的事做出决定。

# 实例

## Colleague

```golang
type Country interface {
   SendMessage(message string)
   GetMessage(message string)
}
type USA struct {
   UnitedNations
}

func (usa USA) SendMessage(message string) {
   usa.UnitedNations.ForwardMessage(message, usa)
}

func (usa USA) GetMessage(message string) {
   fmt.Printf("美国收到对方消息: %s\n", message)
}

type Iraq struct {
   UnitedNations
}

func (iraq Iraq) SendMessage(message string) {
   iraq.UnitedNations.ForwardMessage(message, iraq)
}

func (iraq Iraq) GetMessage(message string) {
   fmt.Printf("伊拉克收到对方消息: %s\n", message)
}
```

## Mediator

```golang
type UnitedNations interface {
   ForwardMessage(message string, country Country)
}

type UnitedNationsSecurityCouncil struct {
   USA
   Iraq
}

func (unsc UnitedNationsSecurityCouncil) ForwardMessage(message string, country Country) {
   switch country.(type) {
   case USA:
      unsc.Iraq.GetMessage(message)
   case Iraq:
      unsc.USA.GetMessage(message)
   default:
      fmt.Println("The country is not a member of UNSC")
   }

}
```

