package builder

import "os"

type HtmlBuilder struct {
	name  string
	fileP *os.File
}

func InitHtmlBuilder(name string) *HtmlBuilder {
	filename := name + ".html"
	var f *os.File
	if _, err := os.Stat(filename); err != nil {
		f, _ = os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	}
	f, _ = os.Create(filename)
	return &HtmlBuilder{name: name, fileP: f}
}

func (h *HtmlBuilder) MakeTitle(s string) {
	h.fileP.WriteString("<html><head><title>" + h.name + "</title></head><body>")
	h.fileP.WriteString("<h1>" + h.name + "</h1>")
}

func (h *HtmlBuilder) MakeString(s string) {
	h.fileP.WriteString("<p>" + s + "</p>")
}

func (h *HtmlBuilder) MakeItems(items ...string) {
	h.fileP.WriteString("<ul>")
	for _, item := range items {
		h.fileP.WriteString("<li>" + item + "</li>")
	}
	h.fileP.WriteString("</ul>")
}

func (h *HtmlBuilder) Close() {
	h.fileP.WriteString("</body></html>")
	defer h.fileP.Close()
}

func (h *HtmlBuilder) GetResult() string {
	return h.name
}
