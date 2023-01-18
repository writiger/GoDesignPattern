package builder

import "strings"

type TextBuilder struct {
	builder strings.Builder
}

func InitTextBuilder() *TextBuilder {
	return &TextBuilder{}
}

func (t *TextBuilder) MakeTitle(s string) {
	t.builder.WriteString("=============================\n")
	t.builder.WriteString("[" + s + "]\n")
	t.builder.WriteString("\n")
}

func (t *TextBuilder) MakeString(s string) {
	t.builder.WriteString("■" + s + "\n")
	t.builder.WriteString("\n")
}

func (t *TextBuilder) MakeItems(items ...string) {
	for _, item := range items {
		t.builder.WriteString("· " + item + "\n")
	}
	t.builder.WriteString("\n")
}

func (t *TextBuilder) Close() {
	t.builder.WriteString("=============================\n")
}

func (t *TextBuilder) GetResult() string {
	return t.builder.String()
}
