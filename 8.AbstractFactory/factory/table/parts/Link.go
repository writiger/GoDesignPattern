package parts

type Link struct {
	url     string
	caption string
}

func (l *Link) MakeHtml() string {
	return `<td><a href="` + l.url + `">` + l.caption + `</a></td>`
}

func InitLink(url, caption string) *Link {
	return &Link{
		url:     url,
		caption: caption,
	}
}
