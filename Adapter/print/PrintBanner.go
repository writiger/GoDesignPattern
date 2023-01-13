package print

type PrintBanner struct {
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
