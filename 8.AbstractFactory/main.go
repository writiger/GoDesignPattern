package main

import (
	"8.AbstractFactory/factory"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: 8.abstract factory 
example: go run main.go list`)
		os.Exit(1)
	}
	f := factory.GetFactory(os.Args[1])
	people := f.CreateLink("人民日报", "https://www.people.com.cn/")
	gmw := f.CreateLink("光明日报", "https://www.gmw.cn")

	bilibili := f.CreateLink("BiliBili!", "https://www.bilibili.com/")
	douyin := f.CreateLink("DouYin", "https://www.douyin.com/")

	bing := f.CreateLink("Bing", "https://cn.bing.com/")
	baidu := f.CreateLink("百度", "https://www.baidu.com/")
	sougou := f.CreateLink("搜狗", "https://www.sogou.com/")

	videos := f.CreateTray("视频")
	videos.Add(bilibili)
	videos.Add(douyin)

	news := f.CreateTray("日报")
	news.Add(people)
	news.Add(gmw)
	news.Add(videos)

	search := f.CreateTray("搜索引擎")
	search.Add(bing)
	search.Add(baidu)
	search.Add(sougou)

	page := f.CreatePage("LinkPage", "Writiger")
	page.Add(news)
	page.Add(search)
	page.Output()
}
