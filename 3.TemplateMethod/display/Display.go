package display

import _interface "TemplateMethod/interface"

// Display 模板
func Display(d _interface.AbstractDisplay) {
	d.Open()
	for i := 0; i < 5; i++ {
		d.Print()
	}
	d.Close()
}
