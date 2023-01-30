# Facade模式

​	程序这东西总会越来越大。随着时间的推移，程序中的类越来越多，而且他们之间相互关联，这会导致程序的结构也越来越复杂。我们在使用这些类之前，必须先弄清楚他们之间的关系，注意正确的调用顺序。

​	特别是处理大型程序时，我们需要格外注意那些数量庞大的类之间的错综复杂的关系。不过与其这么做，不如为这个大型程序准备一个”窗口“。这样，我们就不必单独的关注每个类了。

​	使用Facade模式可以为相互关联在一起的类整理出高层接口（API）。其中的Facade角色可以让系统只有一个简单的接口（API）。而且，Facade角色还会考虑到系统内部的各个类的依赖关系和责任关系，按照正确的顺序调用各个类。

# 实例

## 模拟数据 database.txt

~~~ txt
zhangsan@qq.com=Zhang San
lisi@qq.com=Li Si
wangwu@qq.com=Wang Wu
~~~

## Database结构体

```golang
type database struct {
}

func (d database) GetProperty(email string) (string, error) {
   path := "database.txt"
   file, err := os.OpenFile(path, os.O_RDONLY, 0666)

   if err != nil {
      return "", err
   }

   defer file.Close()

   reader := bufio.NewReader(file)
   for {
      line, _, err := reader.ReadLine()
      if err == io.EOF {
         fmt.Println("File Read Over.")
         break
      }
      info := strings.Split(string(line), "=")
      if email == info[0] {
         return info[1], nil
      }
   }
   return "", errors.New("查询邮箱为空")
}
```

## HtmlWriter结构体

```golang
type htmlMaker struct {
   res string
}

func (hm *htmlMaker) MakeHtml(filename string) error {
   var f *os.File
   f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
   if err != nil {
      return err
   }
   defer f.Close()
   _, err = f.WriteString(hm.res)
   if err != nil {
      return err
   }
   return nil
}

func (hm *htmlMaker) Title(title string) {
   sb := strings.Builder{}
   sb.WriteString("<html><head><title>")
   sb.WriteString(title)
   sb.WriteString("</title></head><body><h1>")
   sb.WriteString(title)
   sb.WriteString("</h1>\n")
   hm.res += sb.String()
}

func (hm *htmlMaker) Paragraph(msg string) {
   hm.res += fmt.Sprintf("<p>%s</p>\n", msg)
}

func (hm *htmlMaker) Link(href, caption string) {
   hm.Paragraph(fmt.Sprintf("<a href=%s>%s</a>", href, caption))
}

func (hm *htmlMaker) MailTo(mail, name string) {
   hm.Link("mailto:"+mail, name)
}

func (hm *htmlMaker) Close() {
   hm.res += "</body></html>\n"
}
```

## PageMaker.go（Facade

```golang
func MakeWelcomePage(mail, filename string) error {
   d := &database{}
   username, err := d.GetProperty(mail)
   if err != nil {
      return err
   }
   hm := &htmlMaker{}
   hm.Title("Welcome to " + username + "'s Page!")
   hm.Paragraph("欢迎来到" + username + "的主页")
   hm.Paragraph("等着你的邮件哦 ！")
   hm.MailTo(mail, username)
   hm.Close()
   if err = hm.MakeHtml(filename); err != nil {
      return err
   }
   fmt.Printf("%s is created for %s ( %s )", filename, mail, username)
   return nil
}
```

# 拓展思路

## Facade角色到底做什么工作

* 使复杂的东西**看起来**简单

* 减少与外部的关联性

## 递归的使用Facade

## 开发人员不愿意使用Facade模式的原因——心理原因

* 当某个程序员得意地说出”啊，在调用那个类之前需要先调用这个类。在调用那个方法前要在这个类中注册一下“的时候，就意味着我们需要引入Facade角色了。