package list

import (
	"8.AbstractFactory/factory/list/parts"
	"8.AbstractFactory/interface/abstract_parts"
)

type Factory struct {
}

func (f Factory) CreateLink(url, caption string) abstract_parts.Link {
	return parts.InitLink(url, caption)
}
func (f Factory) CreateTray(caption string) abstract_parts.Tray {
	return parts.InitTray(caption)
}
func (f Factory) CreatePage(title, name string) abstract_parts.Page {
	return parts.InitPage(title, name)
}
