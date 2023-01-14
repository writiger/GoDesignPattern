package print

import _interface "Adapter/interface"

type PrintBanner struct {
	_interface.Print
	Banner Banner
}

func InitPrintBanner(info string) PrintBanner {
	return PrintBanner{
		Banner: Banner{Info: info},
	}
}

func (pb *PrintBanner) PrintWeak() {
	pb.Banner.ShowWithParen()
}

func (pb *PrintBanner) PrintStrong() {
	pb.Banner.ShowWithAster()
}
