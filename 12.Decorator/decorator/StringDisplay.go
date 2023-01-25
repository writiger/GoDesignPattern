package decorator

type StringDisplay struct {
	info string
}

func InitStringDisplay(info string) *StringDisplay {
	return &StringDisplay{info: info}
}

func (sd StringDisplay) GetColumns() int {
	return len(sd.info)
}

func (sd StringDisplay) GetRows() int {
	return 1
}

func (sd StringDisplay) GetRowText(row int) string {
	if row != 0 {
		return ""
	}
	return sd.info
}
