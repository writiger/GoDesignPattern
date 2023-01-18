package builder

import _interface "7.Builder/interface"

type Director struct {
	builder _interface.Builder
}

func InitDirector(b _interface.Builder) *Director {
	return &Director{builder: b}
}

func (d *Director) Construct() {
	d.builder.MakeTitle("Greeting")
	d.builder.MakeString(" 从早上到下午")
	d.builder.MakeItems("早上好", "下午好")
	d.builder.MakeString("晚上")
	d.builder.MakeItems("晚上好", "晚安", "再见")
	d.builder.Close()
}
