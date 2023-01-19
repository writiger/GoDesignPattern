package parts

import (
	_interface "8.AbstractFactory/interface"
	"os"
	"strings"
)

type Page struct {
	title   string
	author  string
	content []_interface.Item
}

func InitPage(title, name string) *Page {
	return &Page{
		title:  title,
		author: name,
	}
}

func (p *Page) Add(item _interface.Item) {
	p.content = append(p.content, item)
}
func (p *Page) Output() {
	filename := "Page.html"

	var f *os.File
	f, _ = os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	defer f.Close()
	f.WriteString(p.MakeHtml())
}

func (p *Page) MakeHtml() string {
	sb := strings.Builder{}
	sb.WriteString("<html><head><title>" + p.title + "</title></head>\n")
	sb.WriteString("<body>\n")
	sb.WriteString("<h1>" + p.title + "</h1>")

	sb.WriteString(`<table width="80%" border="3">`)
	for _, item := range p.content {
		sb.WriteString("<tr>" + item.MakeHtml() + "</tr>")
	}
	sb.WriteString("</table>\n")
	sb.WriteString("<hr><address>" + p.author + "</address>")
	sb.WriteString("</body></html>\n")
	return sb.String()
}
