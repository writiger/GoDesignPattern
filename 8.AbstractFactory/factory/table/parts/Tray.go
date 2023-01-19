package parts

import (
	_interface "8.AbstractFactory/interface"
	"strconv"
	"strings"
)

type Tray struct {
	caption string
	trays   []_interface.Item
}

func InitTray(caption string) *Tray {
	return &Tray{
		caption: caption,
	}
}

func (t *Tray) Add(item _interface.Item) {
	t.trays = append(t.trays, item)
}

func (t *Tray) MakeHtml() string {
	sb := strings.Builder{}
	sb.WriteString(`<td><table width="100%" border="1"><tr>`)
	sb.WriteString(`<td bgcolor="#cccccc" align="center" colspan="` + strconv.Itoa(len(t.trays)) + `"><b>` + t.caption + `</b></td>`)
	sb.WriteString(`</tr><tr>`)
	for _, item := range t.trays {
		sb.WriteString(item.MakeHtml())
	}
	sb.WriteString("</tr></table></td>")
	return sb.String()
}
