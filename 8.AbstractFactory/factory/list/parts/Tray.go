package parts

import (
	_interface "8.AbstractFactory/interface"
	"strings"
)

type Tray struct {
	caption string
	trays   []_interface.Item
}

func InitTray(caption string) *Tray {
	return &Tray{caption: caption}
}

func (t *Tray) Add(item _interface.Item) {
	t.trays = append(t.trays, item)
}

func (t *Tray) MakeHtml() string {
	sb := strings.Builder{}
	sb.WriteString("<li>\n")
	sb.WriteString(t.caption + "\n")
	sb.WriteString("<ul>\n")
	for _, item := range t.trays {
		sb.WriteString(item.MakeHtml())
	}
	sb.WriteString("</ul>\n</li>\n")
	res := sb.String()
	return res
}
