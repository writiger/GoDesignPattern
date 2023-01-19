package parts

type Link struct {
	url     string
	caption string
}

func (l *Link) MakeHtml() string {
	return `<li><a href="` + l.url + `">` + l.caption + `</a></li>`
}

func InitLink(url, caption string) *Link {
	return &Link{
		url:     url,
		caption: caption,
	}
}
