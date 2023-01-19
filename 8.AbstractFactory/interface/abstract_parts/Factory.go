package abstract_parts

type Factory interface {
	CreateLink(string, string) Link
	CreateTray(string) Tray
	CreatePage(string, string) Page
}
