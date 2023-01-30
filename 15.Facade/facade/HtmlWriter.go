package facade

import (
	"fmt"
	"os"
	"strings"
)

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
